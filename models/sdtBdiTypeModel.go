package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type SdtBdiType struct {
	Id       int    `form:"id"`       //主键
	TypeName string `form:"typeName"` //指标名称
	Remarks  string `form:"typeNameRemarks"`  //指标集备注
}

func (u *SdtBdiType) TableName() string {
	return "sdt_bdi_type"
}

func (this *SdtBdiType) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertSql = " insert into sdt_bdi_type(type_name, remarks) values (?, ?)"
	_, err := o.Raw(insertSql, this.TypeName, this.Remarks).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

func (this *SdtBdiType) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql = "update sdt_bdi " +
		"set  " +
		" type_name = ?, " +
		" remarks = ? " +
		"where id = ?"
	_, err := o.Raw(updateSql, this.TypeName, this.Remarks, this.Id).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}
