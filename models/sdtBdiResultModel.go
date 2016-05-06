package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type SdtBdiResult struct {
	Id         int    `form:"id"`         //主键
	BdiId      int    `form:"bdiId"`      //指标Id
	TableName  string `form:"tableName"`  //表名
	TableLabel string `form:"tableLabel"` //中文表名
	BdiTypeId  int    `form:"bdiTypeId"`  //指标分类ID

	UserCode   string    `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//以下字段为datagrid展示
	BdiTypeName string
	BdiName     string
}

func (u *SdtBdiResult) GetTableName() string {
	return "sdt_bdi_result"
}

/**
获取所有指标。
*/
func (this *SdtBdiResult) GetAllSdtBdiResult(rows int, page int) ([]SdtBdiResult, int, error) {
	var err error
	o := orm.NewOrm()
	sdtBdiResultSlice := make([]SdtBdiResult, 0)

	var querySql = "select r.*, t.type_name as bdi_type_name, bdi.bdi_name as bdi_name from sdt_bdi_result r  " +
		"	left join sdt_bdi_type t on r.bdi_type_id = t.id  " +
		"	left join sdt_bdi bdi on bdi.id = r.bdi_id " +
		"	where r.bdi_id = ? limit ?, ? "

	_, err = o.Raw(querySql, this.BdiId, (page-1)*rows, page*rows).QueryRows(&sdtBdiResultSlice)
	if err != nil {
		log.Fatal("查询表：" + this.GetTableName() + "出错！")
		return nil, 0, err
	}

	num := new(int)
	var countSql = "select count(*) as counts from sdt_bdi_result "

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
func (this *SdtBdiResult) GetSdtBdiResultById() error {
	var o orm.Ormer
	o = orm.NewOrm()

	var querySql = " select r.* from sdt_bdi_result r where r.id = ? "

	err := o.Raw(querySql, this.Id).QueryRow(this)
	if err != nil {
		log.Fatal("查询表：" + this.GetTableName() + "出错！")
		return err
	}

	return nil
}

func (this *SdtBdiResult) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertSdtBdiResultSql = " insert into sdt_bdi_result ( " +
		"	bdi_id, " +
		"	table_name, " +
		"	table_label, " +
		"	bdi_type_id, " +
		"	user_code, " +
		"	create_time " +
		") " +
		" values (?, ?, ?, ?, ?, ?) "

	_, err := o.Raw(insertSdtBdiResultSql, this.BdiId, this.TableName, this.TableLabel, this.BdiTypeId, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

func (this *SdtBdiResult) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSdtBdiResultSql = " update sdt_bdi_result " +
		" set table_name = ?, " +
		" table_label = ?, " +
		" user_code = ?, " +
		" edit_time = ? " +
		" where id = ?"

	_, err := o.Raw(updateSdtBdiResultSql, this.TableName, this.TableLabel, 0, time.Now(), this.Id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}
