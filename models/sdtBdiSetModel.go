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
	BdiSetId     int       `form:"bdiSetId"`   //主键
	BdiSetName   string    `form:"bdiSetName"` //指标集名称
	Remarks      string    `form:"remarks"`    //指标集备注
	CreateUserId int       //创建人ID
	CreateTime   time.Time //创建时间
	ModifyUserId int       //修改人ID
	ModifyTime   time.Time //修改时间

	//以下字段为datagrid展示
	BdiDomainId   string `form:"bdiDomainIds"` //指标域ids
	BdiDomainName string //指标域名称
}

func (u *SdtBdiSet) TableName() string {
	return "sdt_bdi_set"
}

/**
获取所有指标集。
*/
func GetAllSdtBdiSet(rows int, page int) ([]SdtBdiSet, int, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	var sdtBdiSetSlice []SdtBdiSet

	//查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	var querySql = " select " +
		"	s.bdi_set_id, " +
		"	s.bdi_set_name, " +
		"	s.remarks, " +
		"	s.create_user_id, " +
		"	s.create_time, " +
		"	s.modify_user_id, " +
		"	s.modify_time, " +
		"	group_concat(d.bdi_domain_id) as bdi_domain_id, " +
		"	group_concat(d.bdi_domain_name) as bdi_domain_name " +
		"from " +
		"	sdt_bdi_set s, " +
		"	sdt_bdi_domain_set ds, " +
		"	sdt_bdi_domain d " +
		"where " +
		"	s.bdi_set_id = ds.bdi_set_id " +
		"and ds.bdi_domain_id = d.bdi_domain_id " +
		"group by s.bdi_set_id limit ?, ? "

	num, err := o.Raw(querySql, (page-1)*rows, page*rows).QueryRows(&sdtBdiSetSlice)
	if err != nil {
		log.Fatal("查询表：sdt_bdi_set出错！")
		return nil, 0, err
	}

	var countSql = " select count(*) as counts from sdt_bdi_set "
	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：sdt_bdi_set出错！")
		return nil, 0, err
	}

	return sdtBdiSetSlice, int(num), nil
}

/**
获取所有指标集。
*/
func GetSdtBdiSetById(bdiSetId int) (SdtBdiSet, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	var sdtBdiSet SdtBdiSet

	//查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	var querySql = " select " +
		"	s.bdi_set_id, " +
		"	s.bdi_set_name, " +
		"	s.remarks, " +
		"	s.create_user_id, " +
		"	s.create_time, " +
		"	s.modify_user_id, " +
		"	s.modify_time, " +
		"	group_concat(d.bdi_domain_id) as bdi_domain_id, " +
		"	group_concat(d.bdi_domain_name) as bdi_domain_name " +
		" from " +
		"	sdt_bdi_set s, " +
		"	sdt_bdi_domain_set ds, " +
		"	sdt_bdi_domain d " +
		" where " +
		"	s.bdi_set_id = ds.bdi_set_id " +
		" and ds.bdi_domain_id = d.bdi_domain_id " +
		" and s.bdi_set_id = ? "

	err := o.Raw(querySql, bdiSetId).QueryRow(&sdtBdiSet)
	if err != nil {
		log.Fatal("查询表：sdt_bdi_set出错！")
		return sdtBdiSet, err
	}

	return sdtBdiSet, nil
}

func (this *SdtBdiSet)Add() error {
	o := orm.NewOrm()
	o.Begin()

	bdiDomainIds := strings.Split(this.BdiDomainId, ",")

	fmt.Println("bdiDomainIds: ", bdiDomainIds)
	var insertSdtBdiSetSql = " insert into sdt_bdi_set (bdi_set_name, remarks, create_user_id, create_time)values(?, ?, ?, ?)  "
	var insertSdtBdiDomainSetSql = " insert into sdt_bdi_domain_set(bdi_domain_id, bdi_set_id) values (? , ?)   "

	res, err := o.Raw(insertSdtBdiSetSql, this.BdiSetName, this.Remarks, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	bidSetId, err := res.LastInsertId()

	fmt.Println("bidSetId: ", bidSetId)
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
	fmt.Println("this.BdiSetId: ", this.BdiSetId)
	o := orm.NewOrm()
	o.Begin()

	var deleteSdtBdiDomainSetSql = " delete from sdt_bdi_domain_set where bdi_set_id = ? "
	var updateSdtBdiSetSql = " update sdt_bdi_set set  " +
		" bdi_set_name = ?, " +
		" remarks = ?, " +
		//	" modify_user_id = ?, " +
		" modify_time = ? " +
		" where bdi_set_id = ? "
	var insertSdtBdiDomainSetSql = " insert into sdt_bdi_domain_set(bdi_domain_id, bdi_set_id) values (? , ?)   "

	_, err := o.Raw(updateSdtBdiSetSql, this.BdiSetName, this.Remarks, time.Now(), this.BdiSetId).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	_, err = o.Raw(deleteSdtBdiDomainSetSql, this.BdiSetId).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	bdiDomainIds := strings.Split(this.BdiDomainId, ",")
	fmt.Println("bdiDomainIds: ", bdiDomainIds)

	for _, v := range bdiDomainIds {
		bdiDomainId, err := strconv.Atoi(v)
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.Raw(insertSdtBdiDomainSetSql, bdiDomainId, this.BdiSetId).Exec()
		if err != nil {
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
func (this *SdtBdiSet)GetAllSdtBdiSetForList() ([]orm.Params, error) {
	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select bdi_set_id as Id, bdi_set_name as Text from sdt_bdi_set ").Values(&maps)

	if err != nil {
		return nil, err
	}

	return maps, nil
}