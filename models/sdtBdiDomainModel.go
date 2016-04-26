package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type SdtBdiDomain struct {
	Id            int    `form:"id"`            //主键，自动增长
	BdiBaseId     int    `form:"bdiBaseId"`     //指标库Id
	BdiDomainName string `form:"bdiDomainName"` //指标域名称
	AdtFlag       int    `form:"adtFlag"`       //是否使用: 0 - 停用， 1 - 使用
	Remarks       string `form:"remarks"`       //指标库备注

	UserCode   int       `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//以下字段为datagrid展示
	BdiBaseName string //指标库名称
}

func (u *SdtBdiDomain) TableName() string {
	return "sdt_bdi_domain"
}

/**
获取所有指标域。
*/
func (this *SdtBdiDomain) GetAllSdtBdiDomain(rows int, page int) ([]SdtBdiDomain, int, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	var sdtBdiDomainList []SdtBdiDomain = make([]SdtBdiDomain, 0)

	//查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	var querySql = "select " +
		"	sdt_bdi_domain.id, " +
		"	sdt_bdi_domain.bdi_base_id, " +
		"	sdt_bdi_base.bdi_base_name, " +
		"	sdt_bdi_domain.bdi_domain_name, " +
		"	sdt_bdi_domain.adt_flag, " +
		"	sdt_bdi_domain.remarks, " +
		"	sdt_bdi_domain.user_code, " +
		"	sdt_bdi_domain.create_time, " +
		"	sdt_bdi_domain.edit_time " +
		"from 	sdt_bdi_domain " +
		"left join sdt_bdi_base on sdt_bdi_domain.bdi_base_id = sdt_bdi_base.id " + //关联sdt_bdi_base表
		"limit ?, ?"

	num, err := o.Raw(querySql, (page-1)*rows, page*rows).QueryRows(&sdtBdiDomainList)
	if err != nil {
		fmt.Println(err)
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	var countSql = " select count(*) as counts from " + this.TableName()
	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return sdtBdiDomainList, int(num), nil
}

func (this *SdtBdiDomain) Add() error {
	var o orm.Ormer
	o = orm.NewOrm()

	var insertSql = " insert into sdt_bdi_domain " +
		" (bdi_base_id , bdi_domain_name , adt_flag , remarks, create_time) values (?, ?, ?, ?, ?) "

	_, err := o.Raw(insertSql, this.BdiBaseId, this.BdiDomainName, this.AdtFlag, this.Remarks, time.Now()).Exec()
	if err != nil {
		return err
	}

	return nil
}

//更新
func (this *SdtBdiDomain) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql string = " update sdt_bdi_domain " +
		" set bdi_base_id = ?, " +
		" bdi_domain_name = ?, " +
		" adt_flag = ?, " +
		" remarks = ?, " +
		" edit_time = ? " +
		" where id = ? "

	_, err := o.Raw(updateSql, this.BdiBaseId, this.BdiDomainName, this.AdtFlag, this.Remarks, time.Now(), this.Id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

func (this *SdtBdiDomain) GetSdtBdiDomainById() error {
	var o orm.Ormer
	o = orm.NewOrm()

	//查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	var querySql = " select bdi_base_id, bdi_domain_name, adt_flag, remarks " +
		" from sdt_bdi_domain where id = ? "

	err := o.Raw(querySql, this.Id).QueryRow(this)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return err
	}

	return nil
}
