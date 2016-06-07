package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	// "strconv"
	// "strings"
	"database/sql"
	"fmt"
	"strings"
	"time"
	"strconv"
)

type SdtBdiBusiFields struct {
	Id          int       `form:"id"`          //主键
	BusiId      int       `form:"busiId"`      //业务表ID
	Name        string    `form:"name"`        //字段名称
	AsName      string    `form:"asName"`      //字段别名
	Sequence    int       `form:"sequence"`    //字段序号
	Comment     string    `form:"comment"`     //字段注释
	DataType    string    `form:"dataType"`    //字段类型
	DataLength  string    `form:"dataLength"`  //字段长度
	IsProcess   int       `form:"isProcess"`   //是否经过处理
	ProcessType string    `form:"processType"` //处理类型
	Params      string    `form:"params"`      //参数
	UserCode    string    `form:"userCode"`    //用户编码
	CreateTime  time.Time `form:"createTime"`  //创建时间
	EditTime    time.Time `form:"editTime"`    //更新时间

	//以下字段为datagrid展示
	BusiName string
	BdiId    int `form:"bdiId"`
	BdiName  string
	Alias string
}

func (u *SdtBdiBusiFields) TableName() string {
	return "sdt_bdi_busi_fields"
}

/**
 */
func (this *SdtBdiBusiFields) GetAllSdtBdiBusiFields(rows int, page int) ([]SdtBdiBusiFields, int, error) {
	o := orm.NewOrm()
	var sdtBdiBusiFieldsSlice []SdtBdiBusiFields = make([]SdtBdiBusiFields, 0)

	var querySql = " select " +
		"	bdi.id as bdi_id, " +
		"	bdi.bdi_name as bdi_name, " +
		"	busi. name as busi_name, " +
		"	field.* " +
		" from " +
		"	sdt_bdi bdi, " +
		"	sdt_bdi_busi busi, " +
		"	sdt_bdi_busi_fields field " +
		" where " +
		"	bdi.id = ? " +
		" and bdi.id = busi.bdi_id " +
		" and busi.id = field.busi_id " +
		" order by " +
		"	field.sequence " +
		"  	limit ?, ? "

	_, err := o.Raw(querySql, this.BdiId, (page-1)*rows, page*rows).QueryRows(&sdtBdiBusiFieldsSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	num := new(int)
	var countSql = " select " +
		" count(*) as counts " +
		" from " +
		"	sdt_bdi bdi, " +
		"	sdt_bdi_busi busi, " +
		"	sdt_bdi_busi_fields field " +
		" where " +
		"	bdi.id = ? " +
		" and bdi.id = busi.bdi_id " +
		" and busi.id = field.busi_id "

	err = o.Raw(countSql, this.BdiId).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return sdtBdiBusiFieldsSlice, *num, nil
}

/**
 */
func (this *SdtBdiBusiFields) GetSdtBdiBusiById() error {
	o := orm.NewOrm()
	var querySql = " select f.*, b.as_name as alias from sdt_bdi_busi_fields f left join sdt_bdi_busi b on f.busi_id = b.id where f.id = ? "

	err := o.Raw(querySql, this.Id).QueryRow(this)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return err
	}

	return nil
}

func (this *SdtBdiBusiFields) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertSdtBdiBusiSql = "insert into sdt_bdi_busi ( " +
		"	bdi_id, " +
		"	datasource_id, " +
		"	name, " +
		"	user_code, " +
		"	create_time " +
		") " +
		"values (?, ?, ?, ?, ?) "

	_, err := o.Raw(insertSdtBdiBusiSql, this.BdiId, this.Id, this.Name, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (this *SdtBdiBusiFields) AddFields(tableTreeAttributes []TableTreeAttributes) error {
	o := orm.NewOrm()
	o.Begin()

	for _, tableValue := range tableTreeAttributes {
		sdtBdiBusi := SdtBdiBusi{}
		var tableResult sql.Result
		var busiId int

		//先看是否有重复记录
		var sdtBdiBusiSlice []SdtBdiBusi = make([]SdtBdiBusi, 0)
		_, err := o.Raw(" select * from sdt_bdi_busi b where b.bdi_id = ? and b.name = ? ", tableValue.BdiId, tableValue.Name).QueryRows(&sdtBdiBusiSlice)
		if err == nil {
			if len(sdtBdiBusiSlice) > 0 { //有记录
				for _, tempTableValue := range sdtBdiBusiSlice {
					err = o.Raw(" select * from sdt_bdi_busi where id = ? ", tempTableValue.Id).QueryRow(&sdtBdiBusi)
					if err != nil {
						fmt.Println(err)
						o.Rollback()
						return err
					}

					busiId = sdtBdiBusi.Id
				}
			} else { //没有记录
				tableResult, err = o.Raw(" insert into sdt_bdi_busi(bdi_id, datasource_id, name, as_name, cn_name, create_time) values (?, ?, ?, ?, ?, ?)",
					tableValue.BdiId, tableValue.BdiId, tableValue.Name, this.tableAliasName(tableValue.Name), tableValue.CnName, time.Now()).Exec()

				lastInsertId, err := tableResult.LastInsertId()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}

				busiId = int(lastInsertId)
			}
		} else {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		for _, fieldValue := range tableValue.ChildColumns {
			count := new(int)
			var countSql = " select count(*) as counts from sdt_bdi_busi_fields f where  " +
				"	f.busi_id = ? and f.name = ? "

			err = o.Raw(countSql, busiId, fieldValue.Name).QueryRow(&count)
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}

			if *count > 0 {
				continue
			}

			//新增前获取最大 sequence
			num := new(int)
			err = o.Raw(" select max(sequence) from sdt_bdi_busi_fields where busi_id in ( select id from sdt_bdi_busi where bdi_id = ?) ", tableValue.BdiId).QueryRow(&num)
			if err != nil {
				o.Rollback()
				return err
			}

			maxSequence := 0
			if num == nil {
				maxSequence = 0
			} else {
				maxSequence = *num
			}

			var dataLength string
			if fieldValue.DataLength < 1 {
				dataLength = ""
			}else {
				dataLength = strconv.Itoa(fieldValue.DataLength)
			}

			fmt.Println("fieldValue.DataLength: ", fieldValue.DataLength)
			fmt.Println("dataLength: ", dataLength)

			_, err = o.Raw(" insert into sdt_bdi_busi_fields(bdi_id, busi_id, name, sequence, comment, "+
				"data_type, data_length, user_code, create_time) values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
				tableValue.BdiId, busiId,
				fieldValue.Name, maxSequence + 1, fieldValue.Comment, fieldValue.DataType,
				dataLength, 0, time.Now()).Exec()

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

func (this *SdtBdiBusiFields) AddConstFields() error {
	o := orm.NewOrm()
	o.Begin()

	//往 sdt_bdi_busi， 新增一条虚拟表记录
	var insertSdtBdiBusiSql = "insert into sdt_bdi_busi ( " +
		"	bdi_id, " +
		"	datasource_id, " +
		"	name, " +
		"	cn_name, " +
		"	is_process, " +
		"	process_type, " +
		"	params, " +
		"	user_code, " +
		"	create_time " +
		") " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?) "

	res, err := o.Raw(insertSdtBdiBusiSql, this.BdiId, 0, "virtual_table", "虚拟表", 0, "const", "", 0, time.Now()).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	busiId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	num := new(int)
	var selectMaxSequence = " select max(sequence) from sdt_bdi_busi_fields where busi_id in ( select id from sdt_bdi_busi where bdi_id = ?) "
	err = o.Raw(selectMaxSequence, this.BdiId).QueryRow(&num)
	if err != nil {
		o.Rollback()
		return err
	}

	maxSequence := 0
	if num == nil {
		maxSequence = 0
	} else {
		maxSequence = *num
	}

	var insertFieldSql = " insert into sdt_bdi_busi_fields(is_process, busi_id, bdi_id, name, sequence, comment, data_type, data_length, process_type, params, user_code, create_time) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err = o.Raw(insertFieldSql, 1, busiId, this.BdiId, this.Name, maxSequence+1, this.Comment, this.DataType, this.DataLength, "const", this.Params, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

func (this *SdtBdiBusiFields) Update() error {
	//o := orm.NewOrm()
	//o.Begin()
	//
	//var updateSql = "update sdt_bdi_busi " +
	//	" set  " +
	//	" datasource_id = ?, " +
	//	" name = ?, " +
	//	" user_code = ?, " +
	//	" edit_time = ? " +
	//	" where id = ?"
	//
	//_, err := o.Raw(updateSql, this.Id, this.Name, this.UserCode, time.Now(), this.Id).Exec()
	//if err != nil {
	//	o.Rollback()
	//	return err
	//}
	//o.Commit()
	return nil
}

func (this *SdtBdiBusiFields) UpdateFields() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql = " update sdt_bdi_busi_fields set " +
		" process_type = ?, " +
		" params = ?, " +
		" data_type = ?, " +
		" data_length = ?, " +
		" is_process = 1, " +
		" as_name = ?, " +
		" edit_time = ? " +
		//" data_type = ?, " +
		//" data_length = ? " +
		" where id = ? "

	_, err := o.Raw(updateSql, this.ProcessType, this.Params, this.DataType, this.DataLength, this.AsName, time.Now(), /*this.DataType, this.DataLength,*/ this.Id).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (this *SdtBdiBusiFields) AddBusiAndAddField(tableTreeAttributes []TableTreeAttributes) error {
	o := orm.NewOrm()
	o.Begin()

	//num := new(int)
	//var countSql = " select count(*) as counts from sdt_bdi_busi b where b.bdi_id = ? and b.name = ? "
	//err := o.Raw(countSql).QueryRow(num)

	var insertTableSql = " insert into sdt_bdi_busi(bdi_id, datasource_id, name, cn_name, create_time) values (?, ?, ?, ?, ?)"
	var insertFieldSql = " insert into sdt_bdi_busi_fields(busi_id, name, sequence, comment, data_type, data_length, user_code, create_time) values (?, ?, ?, ?, ?, ?, ?, ?)"
	var sequenceIndex int = 1

	for _, tableValue := range tableTreeAttributes {
		//先看是否有重复记录，如果有，先删除记录
		var sdtBdiBusiSlice []SdtBdiBusiFields = make([]SdtBdiBusiFields, 0)
		var querySql = " select * from sdt_bdi_busi b where b.bdi_id = ? and b.name = ? "
		_, err := o.Raw(querySql, tableValue.BdiId, tableValue.Name).QueryRows(&sdtBdiBusiSlice)
		if err == nil {
			if len(sdtBdiBusiSlice) > 0 { //如果有记录，先删除记录
				for _, tempTableValue := range sdtBdiBusiSlice {
					_, err = o.Raw(" delete from sdt_bdi_busi where bdi_id = ? and name = ? ", tempTableValue.Id, tempTableValue.Name).Exec()
					if err != nil {
						fmt.Println(err)
						o.Rollback()
						return err
					}

					_, err = o.Raw(" delete from sdt_bdi_busi_fields where busi_id = ? ", tempTableValue.Id).Exec()
					if err != nil {
						fmt.Println(err)
						o.Rollback()
						return err
					}
				}
			}
		} else {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		tableResult, err := o.Raw(insertTableSql, tableValue.BdiId, tableValue.BdiId, tableValue.Name, tableValue.CnName, time.Now()).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		lastInsertId, err := tableResult.LastInsertId()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		for _, fieldValue := range tableValue.ChildColumns {
			////新增前获取最大RuleSetLevel
			//var querySql string = " select max(rule_set_level) as rule_set_level from sdt_bdi_rule_set where bdi_id = ? "
			//maxLevel := new(int)
			//err = o.Raw(querySql, this.BdiId).QueryRow(maxLevel)
			//
			//if err != nil {
			//	o.Rollback()
			//	return err
			//}

			_, err = o.Raw(insertFieldSql, lastInsertId, fieldValue.Name, sequenceIndex, fieldValue.Comment, fieldValue.DataType,
				fieldValue.DataLength, 0, time.Now()).Exec()

			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}

			sequenceIndex = sequenceIndex + 1
		}
	}

	o.Commit()
	return nil
}

//行上移
func (this *SdtBdiBusiFields) RowMoveUp() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	fmt.Println("bdiId: ", this.BdiId)
	//先查询出 sequence 减 1 的行记录id
	var incFieldsIdSql string = " select f.id from sdt_bdi_busi_fields f where " +
	"	f.bdi_id = (select t.bdi_id from sdt_bdi_busi_fields t where id = ?) " +
	" and f.sequence = (select t.sequence - 1 from sdt_bdi_busi_fields t where id = ?) limit 1 "

	incFieldsId := new(int)
	err = o.Raw(incFieldsIdSql, this.Id, this.Id).QueryRow(incFieldsId)
	if err != nil {
		fmt.Println("err1: ", err)
		o.Rollback()
		return err
	}

	var updateSql string = " update sdt_bdi_busi_fields set sequence = ?, edit_time = ? where id = ? "

	//1. 更新下移行记录的 Sequence
	_, err = o.Raw(updateSql, this.Sequence, time.Now(), incFieldsId).Exec()
	if err != nil {
		fmt.Println("err2: ", err)
		o.Rollback()
		return err
	}

	//2. 更新上移行记录的 Sequence
	_, err = o.Raw(updateSql, this.Sequence-1, time.Now(), this.Id).Exec()
	if err != nil {
		fmt.Println("err3: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//行下移
func (this *SdtBdiBusiFields) RowMoveDown() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	//先查询出 sequence 加 1 的行记录id
	var decFieldsIdSql string =
	" select f.id from sdt_bdi_busi_fields f where " +
	"	f.bdi_id = (select t.bdi_id from sdt_bdi_busi_fields t where id = ?) " +
	" and f.sequence = (select t.sequence + 1 from sdt_bdi_busi_fields t where id = ?) limit 1 "
	decFieldId := new(int)

	err = o.Raw(decFieldsIdSql, this.BdiId, this.Sequence+1).QueryRow(decFieldId)
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	var updateSql string = " update sdt_bdi_busi_fields set sequence = ?, edit_time = ? where id = ? "

	//1. 更新下移行记录的 Sequence
	_, err = o.Raw(updateSql, this.Sequence+1, time.Now(), this.Id).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	//2. 更新上移行记录的 Sequence
	_, err = o.Raw(updateSql, this.Sequence, time.Now(), decFieldId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

func (this *SdtBdiBusiFields) CheckData() (int, error) {
	var err error
	o := orm.NewOrm()

	var querySql = " select " +
		"	count(*) as counts " +
		" from " +
		"	sdt_bdi_busi_fields f, " +
		"	sdt_bdi_busi b " +
		" where " +
		"	b.bdi_id = ? " +
		" and b.id = f.busi_id " +
		" and f.is_process = 0 "

	num := new(int)
	err = o.Raw(querySql, this.BdiId).QueryRow(num)
	if err != nil {
		return 0, nil
	}

	return *num, err
}

func (this *SdtBdiBusiFields) Synchronize() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	var insertString = " insert into sdt_bdi_result_fields ( " +
		"	result_id, " +
		"	name, " +
		"	sequence, " +
		"	comment, " +
		"	data_type, " +
		"	data_length, " +
		"	create_time " +
		" ) " +
		"select t.* " +
		"from ( " +
		"		select  r.id as result_id, " +
		"			if(lower(trim(f.name)) = \"id\", concat(busi.name, \"_\",lower(trim(f.name))), lower(trim(f.name))) as name, " +
		"			f.sequence, " +
		"			f. comment, " +
		"			f.data_type, " +
		"			f.data_length, " +
		"			now() as create_time " +
		"		from " +
		"			sdt_bdi_busi_fields f, " +
		"			sdt_bdi b, " +
		"			sdt_bdi_result r, " +
		"			sdt_bdi_busi busi " +
		"		where " +
		"			f.bdi_id = ? " +
		"		and f.bdi_id = b.id " +
		"		and r.bdi_id = b.id " +
		"		and f.busi_id = busi.id " +
		"		order by f.sequence " +
		"	) as t " +
		"where " +
		"	not exists ( " +
		"		select " +
		"			f.id " +
		"		from " +
		"			sdt_bdi_result_fields f " +
		"		where " +
		"			f.result_id = t.result_id " +
		"		and lower(trim(f. name)) = lower(trim(t. name)) " +
		"	) "
	_, err = o.Raw(insertString, this.BdiId).Exec()
	if err != nil {
		fmt.Println(err)
		fmt.Println("更新数据出错！")
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//给表取别名
func (this *SdtBdiBusiFields) tableAliasName(tableName string) string {
	//给表取别名
	var asName string = ""
	tablesNameArray := strings.Split(tableName, "_")
	for _, v := range tablesNameArray {
		if len(v) >= 1 {
			asName = asName + string(strings.TrimSpace(v)[0])
		}
	}

	return asName
}
