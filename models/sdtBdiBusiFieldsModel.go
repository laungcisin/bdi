package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	// "strconv"
	// "strings"
	"fmt"
	"time"
)

type SdtBdiBusiFields struct {
	Id          int       `form:"id"`          //主键
	BusiId      int       `form:"busiId"`      //业务表ID
	Name        string    `form:"name"`        //字段名称
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
	var querySql = " select * from sdt_bdi_busi_fields busi where busi.id = ? "

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

func (this *SdtBdiBusiFields) AddFields(sdtBdiBusiFieldsSlice []SdtBdiBusiFields) error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	for _, v := range sdtBdiBusiFieldsSlice {
		count := new(int)
		var countSql = " select count(*) as counts from sdt_bdi_busi_fields f where  " +
			"	f.busi_id in (select b.bdi_id from sdt_bdi_busi b where b.bdi_id = ?) " +
			" and f.name = ? "
		err = o.Raw(countSql, v.BdiId, v.Name).QueryRow(&count)
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		if *count > 0 {
			continue
		}

		busiId := new(int)
		var bdiIdQuerrySql = " select busi.id from sdt_bdi_busi busi where busi.id in (select b.bdi_id from sdt_bdi_busi b where b.bdi_id = ?) limit 1 "
		err = o.Raw(bdiIdQuerrySql, v.BdiId).QueryRow(&busiId)
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		num := new(int)
		var selectMaxSequence = " select max(sequence) from sdt_bdi_busi_fields where busi_id in ( select busi_id from sdt_bdi_busi where bdi_id = ?) "
		err = o.Raw(selectMaxSequence, v.BdiId).QueryRow(&num)
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		var insertFieldSql = " insert into sdt_bdi_busi_fields(busi_id, bdi_id, name, sequence, data_type, data_length, user_code, create_time) " +
			"values (?, ?, ?, ?, ?, ?, ?, ?)"

		maxSequence := 0
		if num == nil {
			maxSequence = 0
		} else {
			maxSequence = *num
		}

		_, err = o.Raw(insertFieldSql, busiId, v.BdiId, v.Name, maxSequence+1, v.DataType, v.DataLength, 0, time.Now()).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
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
	var selectMaxSequence = " select max(sequence) from sdt_bdi_busi_fields where busi_id in ( select busi_id from sdt_bdi_busi where bdi_id = ?) "
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

	var insertFieldSql = " insert into sdt_bdi_busi_fields(busi_id, bdi_id, name, sequence, comment, data_type, data_length, process_type, params, user_code, create_time) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err = o.Raw(insertFieldSql, busiId, this.BdiId, this.Name, maxSequence+1, this.Comment, this.DataType, this.DataLength, "const", this.Params, 0, time.Now()).Exec()
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
		" comment = ?, " +
		" edit_time = ? " +
		" where id = ? "

	_, err := o.Raw(updateSql, this.ProcessType, this.Params, this.Comment, time.Now(), this.Id).Exec()
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
			fmt.Println("sequenceIndex: ", sequenceIndex)
		}
	}

	o.Commit()
	return nil
}

//行上移
func (this *SdtBdiBusiFields) RowMoveUp() error {
	fmt.Println("SdtBdiBusiFields: ", this.BdiId, this.Sequence)
	var err error
	o := orm.NewOrm()
	o.Begin()

	//先查询出 sequence 减 1 的行记录id
	var incFieldsIdSql string = " select id from sdt_bdi_busi_fields  " +
		"	where busi_id in( select id from sdt_bdi_busi where bdi_id = ? )  " +
		"	and sequence = ? limit 1"

	incFieldsId := new(int)
	err = o.Raw(incFieldsIdSql, this.BdiId, this.Sequence-1).QueryRow(incFieldsId)
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
	var decFieldsIdSql string = " select id from sdt_bdi_busi_fields  " +
		"	where busi_id in( select id from sdt_bdi_busi where bdi_id = ? )  " +
		"	and sequence = ? limit 1"
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


func (this *SdtBdiBusiFields)Synchronize() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	var insertString =
		" insert into sdt_bdi_result_fields ( " +
		"	result_id, " +
		"	name, " +
		"	sequence, " +
		"	comment, " +
		"	data_type, " +
		"	data_length, " +
		"	create_time " +
		" )select " +
		"	r.id as result_id, " +
		"	lower(f.name) as name, " +
		"	f.sequence, " +
		"	f. comment, " +
		"	f.data_type, " +
		"	f.data_length, " +
		"	now() as create_time " +
		" from " +
		"	sdt_bdi_busi_fields f, " +
		"	sdt_bdi b, " +
		"	sdt_bdi_result r " +
		" where " +
		"	f.bdi_id = ? " +
		" and f.bdi_id = b.id " +
		" and r.bdi_id = b.id order by f.sequence "
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
