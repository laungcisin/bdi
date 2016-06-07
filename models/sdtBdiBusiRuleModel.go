package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"time"
)

type SdtBdiBusiRule struct {
	Id           int    `form:"id"`           //主键
	BdiId        int    `form:"bdiId"`        //指标id
	FieldId      int    `form:"fieldId"`      //业务字段ID
	Name         string `form:"name"`         //规则名称
	OperatorType string `form:"operatorType"` //运算符类型
	Operator     string `form:"operator"`     //运算符
	Params       string `form:"params"`       //参数

	UserCode   string    `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//以下字段为datagrid展示
	BdiName          string
	OperatorTypeName string
	OperatorName     string
}

func (u *SdtBdiBusiRule) TableName() string {
	return "sdt_bdi_busi_rule"
}

/**
获取所有指标。
*/
func (this *SdtBdiBusiRule) GetAllSdtBdiBusiRule(rows int, page int) ([]SdtBdiBusiRule, int, error) {
	o := orm.NewOrm()
	var sdtBdiBusiRuleSlice []SdtBdiBusiRule = make([]SdtBdiBusiRule, 0)

	var querySql = "select " +
		"	t.*, t1.process_name as operator_type_name, " +
		"	t2.process_name as operator_name " +
		"from " +
		"	( " +
		"		select " +
		"			r.*, bdi.bdi_name as bdi_name " +
		"		from " +
		"			sdt_bdi_busi_rule r, " +
		"			sdt_bdi bdi " +
		"		where " +
		"			r.bdi_id = ? " +
		"		and r.bdi_id = bdi.id " +
		"  		limit ?, ? " +
		"	) as t " +
		"left join sdt_bdi_process_type t1 on t1.main_class = 4 " +
		"and t1.sub_class = 1 " +
		"and t1.process_type = t.operator_type " +
		"left join sdt_bdi_process_type t2 on t2.main_class = 3 " +
		"and t2.sub_class = t.operator_type " +
		"and t2.process_type = t.operator"

	_, err := o.Raw(querySql, this.BdiId, (page-1)*rows, page*rows).QueryRows(&sdtBdiBusiRuleSlice)
	if err != nil {
		fmt.Println(err)
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	num := new(int)
	var countSql = "select " +
		"	count(*) as counts " +
		" from " +
		"	sdt_bdi_busi_rule r, " +
		"	sdt_bdi bdi " +
		" where " +
		" 	r.bdi_id = ? " +
		"	and r.bdi_id = bdi.id "

	err = o.Raw(countSql, this.BdiId).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return sdtBdiBusiRuleSlice, *num, nil
}

func (this *SdtBdiBusiRule) GetSdtBdiBusiRuleById() error {
	o := orm.NewOrm()
	var querySql = " select r.* from sdt_bdi_busi_rule r where r.id = ? "

	err := o.Raw(querySql, this.Id).QueryRow(this)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return err
	}

	return nil
}

func (this *SdtBdiBusiRule) AddColumn(bdiId int, fieldsIds []string, tableTreeAttributes []TableTreeAttributes) error {
	var err error
	var latestBusiId int64
	o := orm.NewOrm()
	o.Begin()

	for _, tableValue := range tableTreeAttributes {
		count := new(int)
		err = o.Raw(" select count(*) as count from sdt_bdi_busi b where b.bdi_id = ? and b.name = ? ", tableValue.BdiId, tableValue.Name).QueryRow(&count)
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		if err == nil {
			if *count > 0 { //有记录, 取busiId
				err = o.Raw(" select b.id as count from sdt_bdi_busi b where b.bdi_id = ? and b.name = ? limit 1  ", tableValue.BdiId, tableValue.Name).QueryRow(&latestBusiId)
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			} else { //没有记录，新增一条记录, 取busiId
				tableResult, err := o.Raw(" insert into sdt_bdi_busi(bdi_id, datasource_id, name, cn_name, create_time) values (?, ?, ?, ?, ?) ", tableValue.BdiId, tableValue.DatasourceId, tableValue.Name, tableValue.CnName, time.Now()).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}

				latestBusiId, err = tableResult.LastInsertId()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}

			}
		} else {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		for _, fieldValue := range tableValue.ChildColumns {
			num := new(int)
			var selectMaxSequence = " select max(sequence) from sdt_bdi_busi_fields where busi_id in ( select id from sdt_bdi_busi where bdi_id = ?) "
			err = o.Raw(selectMaxSequence, bdiId).QueryRow(&num)
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}

			maxSequence := 0
			if num == nil {
				maxSequence = 0
			} else {
				maxSequence = *num
			}

			count := new(int)
			err = o.Raw(" select count(*) as counts from sdt_bdi_busi_fields f where f.busi_id = ? and f.bdi_id = ? and f.name = lower(?) ", latestBusiId, bdiId, fieldValue.Name).QueryRow(&count)
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
			var insertFieldSql = " insert into sdt_bdi_busi_fields(busi_id, bdi_id, is_result_column, name, sequence, data_type, data_length, user_code, create_time) " +
				" values (?, ?, 0, ?, ?, ?, ?, ?, ?)"

			var lastFieldsId int64 = 0

			if *count < 1 { //没有，新增
				var dataLength string
				if fieldValue.DataLength < 1 {
					dataLength = ""
				}else {
					dataLength = string(fieldValue.DataLength)
				}
				res, err := o.Raw(insertFieldSql, latestBusiId, bdiId, fieldValue.Name, maxSequence, fieldValue.DataType, dataLength, 0, time.Now()).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}

				lastFieldsId, err = res.LastInsertId()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}

			} else {
				err = o.Raw("  select f.id as counts from sdt_bdi_busi_fields f where f.busi_id = ? and f.bdi_id = ? and f.name = lower(?)  ", latestBusiId, bdiId, fieldValue.Name).QueryRow(&lastFieldsId)
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			}

			fmt.Println("lastFieldsId: ", lastFieldsId)
			if lastFieldsId > 0 {
				_, err = o.Raw(" insert into sdt_bdi_busi_rule(bdi_id, field_id, user_code, create_time)values(?, ?, ?, ?) ", bdiId, lastFieldsId, 0, time.Now()).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}
			}
		}
	}

	if len(fieldsIds) > 0 {
		for _, v := range fieldsIds {
			fieldsId, _ := strconv.Atoi(v)
			if fieldsId < 1 {
				continue
			}

			var insertSql = " insert into sdt_bdi_busi_rule (bdi_id, field_id, name, user_code, create_time) " +
			" select f.bdi_id, f.id, f.comment, 1, now() from sdt_bdi_busi_fields f where f.id = ? limit 1 "
			_, err = o.Raw(insertSql, fieldsId).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}
	}

	o.Commit()
	return nil
}

//更新
func (this *SdtBdiBusiRule) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql string = " update sdt_bdi_busi_rule " +
		" set  " +
		" name = ?, " +
		" operator_type = ?, " +
		" operator = ?, " +
		" params = ?, " +
		//" user_code = '0', " +
		" edit_time = ? " +
		" where id = ? "

	_, err := o.Raw(updateSql, this.Name, this.OperatorType, this.Operator, this.Params, time.Now(), this.Id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}
