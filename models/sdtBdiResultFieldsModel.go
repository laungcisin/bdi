package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type SdtBdiResultFields struct {
	Id         int    `form:"id"`         //主键
	ResultId   int    `form:"resultId"`   //指标结果表ID
	Name       string `form:"name"`       //字段名称
	Sequence   int    `form:"sequence"`   //字段序号
	Comment    string `form:"comment"`    //字段注释
	DataType   string `form:"dataType"`   //字段类型
	DataLength string `form:"dataLength"` //字段长度

	UserCode   string    `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//以下字段为datagrid展示
	TableName string
}

func (u *SdtBdiResultFields) GetTableName() string {
	return "std_bdi_result_fields"
}

func (this *SdtBdiResultFields) GetAllSdtBdiResultFields(rows int, page int) ([]SdtBdiResultFields, int, error) {
	var err error
	o := orm.NewOrm()
	sdtBdiResultSlice := make([]SdtBdiResultFields, 0)

	var querySql = " select f.*, r.table_name from sdt_bdi_result_fields f, sdt_bdi_result r where r.id = ? and f.result_id = r.id limit ?, ? "

	_, err = o.Raw(querySql, this.ResultId, (page-1)*rows, page*rows).QueryRows(&sdtBdiResultSlice)
	if err != nil {
		log.Fatal("查询表：" + this.GetTableName() + "出错！")
		return nil, 0, err
	}

	num := new(int)
	var countSql = " select count(*) as counts from sdt_bdi_result_fields f, sdt_bdi_result r where f.result_id = r.id  "

	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：" + this.GetTableName() + "出错！")
		return nil, 0, err
	}

	return sdtBdiResultSlice, *num, nil
}

/**
获取所有指标。
*/
func (this *SdtBdiResultFields) GetSdtBdiResultFieldsById() error {
	// var o orm.Ormer
	// o = orm.NewOrm()

	// var querySql = " select r.* from std_bdi_result r where r.id = ? "

	// err := o.Raw(querySql, this.Id).QueryRow(this)
	// if err != nil {
	// 	log.Fatal("查询表：" + this.GetTableName() + "出错！")
	// 	return err
	// }

	return nil
}

func (this *SdtBdiResultFields) Add() error {
	// o := orm.NewOrm()
	// o.Begin()

	// var insertSdtBdiResultSql = " insert into std_bdi_result ( " +
	// 	"	table_name, " +
	// 	"	table_label, " +
	// 	"	bdi_type_id, " +
	// 	"	user_code, " +
	// 	"	create_time " +
	// 	") " +
	// 	" values (?, ?, ?, ?, ?) "

	// _, err := o.Raw(insertSdtBdiResultSql, this.TableName, this.TableLabel, this.BdiTypeId, 0, time.Now()).Exec()
	// if err != nil {
	// 	o.Rollback()
	// 	return err
	// }

	// o.Commit()
	return nil
}

func (this *SdtBdiResultFields) Update() error {
	// o := orm.NewOrm()
	// o.Begin()

	// var updateSdtBdiResultSql = " update std_bdi_result " +
	// 	" set table_name = ?, " +
	// 	" table_label = ?, " +
	// 	" bdi_type_id = ?, " +
	// 	" user_code = ?, " +
	// 	" edit_time = ? " +
	// 	" where id = ?"

	// _, err := o.Raw(updateSdtBdiResultSql, this.TableName, this.TableLabel, this.BdiTypeId, 0, time.Now(), this.Id).Exec()
	// if err != nil {
	// 	o.Rollback()
	// 	return err
	// }

	// o.Commit()
	return nil
}
