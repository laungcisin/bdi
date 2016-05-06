package models

import (
	//	"fmt"
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"strings"
	"time"
)

type SdtBdiSet struct {
	Id         int    `form:"id"`         //主键
	BdiSetName string `form:"bdiSetName"` //指标集名称
	Remarks    string `form:"remarks"`    //指标集备注

	UserCode   string    `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//以下字段为datagrid展示
	BdiDomainId   string `form:"sdtBdiDomainIds"` //指标域ids
	BdiDomainName string //指标域名称
}

func (this *SdtBdiSet) TableName() string {
	return "sdt_bdi_set"
}

/**
获取所有指标集。
*/
func (this *SdtBdiSet) GetAllSdtBdiSet(rows int, page int) ([]SdtBdiSet, int, error) {
	o := orm.NewOrm()
	var sdtBdiSetSlice []SdtBdiSet = make([]SdtBdiSet, 0)

	var querySql = " select " +
		"	s.id, " +
		"	s.bdi_set_name, " +
		"	s.remarks, " +
		"	s.user_code, " +
		"	s.create_time, " +
		"	s.edit_time, " +
		"	group_concat(d.id) as bdi_domain_id, " +
		"	group_concat(d.bdi_domain_name) as bdi_domain_name " +
		"from " +
		"	sdt_bdi_set s, " +
		"	sdt_bdi_domain_set rel, " +
		"	sdt_bdi_domain d " +
		"where " +
		"	s.id = rel.set_id " +
		"and rel.domain_id = d.id " +
		"group by s.id limit ?, ? "

	num, err := o.Raw(querySql, (page-1)*rows, page*rows).QueryRows(&sdtBdiSetSlice)
	if err != nil {
		fmt.Println(err)
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	var countSql string = "select " +
		"	count(*) as counts " +
		"from " +
		"	sdt_bdi_set s, " +
		"	sdt_bdi_domain_set rel, " +
		"	sdt_bdi_domain d " +
		"where " +
		"	s.id = rel.set_id " +
		"and rel.domain_id = d.id"

	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		fmt.Println(err)
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return sdtBdiSetSlice, int(num), nil
}

func (this *SdtBdiSet) GetSdtBdiSetById() error {
	o := orm.NewOrm()

	//查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	var querySql = " select " +
		"	s.id, " +
		"	s.bdi_set_name, " +
		"	s.remarks, " +
		"	s.user_code, " +
		"	s.create_time, " +
		"	s.edit_time, " +
		"	group_concat(d.id) as bdi_domain_id, " +
		"	group_concat(d.bdi_domain_name) as bdi_domain_name " +
		" from " +
		"	sdt_bdi_set s, " +
		"	sdt_bdi_domain_set rel, " +
		"	sdt_bdi_domain d " +
		" where " +
		"	s.id = rel.set_id " +
		" and rel.domain_id = d.id " +
		" and s.id = ? "

	err := o.Raw(querySql, this.Id).QueryRow(this)
	if err != nil {
		log.Fatal("查询表：sdt_bdi_set出错！")
		return err
	}

	return nil
}

func (this *SdtBdiSet) Add() error {
	o := orm.NewOrm()
	o.Begin()

	bdiDomainIds := strings.Split(this.BdiDomainId, ",")

	var insertSdtBdiSetSql = " insert into sdt_bdi_set (bdi_set_name, remarks, user_code, create_time)values(?, ?, ?, ?)  "
	var insertSdtBdiDomainSetSql = " insert into sdt_bdi_domain_set(domain_id, set_id) values (? , ?)   "

	res, err := o.Raw(insertSdtBdiSetSql, this.BdiSetName, this.Remarks, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	bidSetId, err := res.LastInsertId()
	if err != nil {
		o.Rollback()
		return err
	}

	for _, v := range bdiDomainIds {
		bdiDomainId, err := strconv.Atoi(v)
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.Raw(insertSdtBdiDomainSetSql, bdiDomainId, bidSetId).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	o.Commit()
	return nil
}

func (this *SdtBdiSet) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var deleteSdtBdiDomainSetSql = " delete from sdt_bdi_domain_set where set_id = ? "
	var updateSdtBdiSetSql = " update sdt_bdi_set " +
		" set bdi_set_name = ?, " +
		" remarks = ?, " +
		//	" modify_user_id = ?, " +
		" edit_time = ? " +
		" where id = ? "
	var insertSdtBdiDomainSetSql = " insert into sdt_bdi_domain_set(domain_id, set_id) values (? , ?)   "

	_, err := o.Raw(updateSdtBdiSetSql, this.BdiSetName, this.Remarks, time.Now(), this.Id).Exec()
	if err != nil {
		fmt.Println("err1: ", err)
		o.Rollback()
		return err
	}

	_, err = o.Raw(deleteSdtBdiDomainSetSql, this.Id).Exec()
	if err != nil {
		fmt.Println("err2: ", err)
		o.Rollback()
		return err
	}

	bdiDomainIds := strings.Split(this.BdiDomainId, ",")

	for _, v := range bdiDomainIds {
		bdiDomainId, err := strconv.Atoi(v)
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.Raw(insertSdtBdiDomainSetSql, bdiDomainId, this.Id).Exec()
		if err != nil {
			fmt.Println("err3: ", err)
			o.Rollback()
			return err
		}
	}

	o.Commit()
	return nil
}

/**
获取所有指标集-用于下拉列表。
*/
func (this *SdtBdiSet) GetAllSdtBdiSetForList() ([]orm.Params, error) {
	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select id as Id, bdi_set_name as Text from sdt_bdi_set ").Values(&maps)

	if err != nil {
		return nil, err
	}

	return maps, nil
}
