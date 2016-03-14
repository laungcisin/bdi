package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"strings"
	"time"
)

type SdtBdi struct {
	BdiId        int    `form:"bdiId"`        //主键
	BdiName      string `form:"bdiName"`      //指标名称
	BdiSecrecyId int    `form:"bdiSecrecyId"` //指标密级ID
	BdiTypeId    int    `form:"bdiTypeId"`    //指标类型ID

	Remarks string `form:"remarks"` //指标集备注

	CreateUserId int       //创建人ID
	CreateTime   time.Time //创建时间
	ModifyUserId int       //修改人ID
	ModifyTime   time.Time //修改时间

	//以下字段为datagrid展示
	BdiTypeName    string //所属指标类型名称
	BdiSecrecyName string //所属指标密级级别名称
	BdiSetIds      string `form:"bdiSetIds"` //所属指标集ids
	BdiSetNames    string //所属指标集名称
}

func (u *SdtBdi) TableName() string {
	return "sdt_bdi"
}

/**
获取所有指标。
*/
func (this *SdtBdi) GetAllSdtBdi(rows int, page int) ([]SdtBdi, int, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	var sdtBdiSlice []SdtBdi

	//查询的字段顺序最好和model的字段对应，方便解析并赋值。
	var querySql = "select " +
		"	sb.bdi_id, " +
		"	sb.bdi_name, " +
		"	sb.bdi_type_id, " +
		"	sb.bdi_secrecy_id, " +
		"	sb.remarks, " +
		"	sbt.bdi_type_name, " +
		"	sbsec.bdi_secrecy_name, " +
		"	group_concat(sbset.bdi_set_id) as bdi_set_ids, " +
		"	group_concat(sbset.bdi_set_name) as bdi_set_names " +
		"from " +
		"	sdt_bdi sb, " +
		"	sdt_bdi_set_rel_bdi rel, " +
		"	sdt_bdi_set sbset, " +
		"	sdt_bdi_type sbt, " +
		"	sdt_bdi_secrecy sbsec " +
		"where " +
		"	sb.bdi_id = rel.bdi_id " +
		"and rel.bdi_set_id = sbset.bdi_set_id " +
		"and sb.bdi_type_id = sbt.bdi_type_id " +
		"and sb.bdi_secrecy_id = sbsec.bdi_secrecy_id " +
		"group by sb.bdi_id limit ?, ? "

	num, err := o.Raw(querySql, (page-1)*rows, page*rows).QueryRows(&sdtBdiSlice)
	if err != nil {
		log.Fatal("查询表：sdt_bdi出错！")
		return nil, 0, err
	}

	var countSql = " select count(*) as counts from " + this.TableName()
	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：sdt_bdi出错！")
		return nil, 0, err
	}

	return sdtBdiSlice, int(num), nil
}

/**
获取所有指标。
*/
func (this *SdtBdi) GetSdtBdiById() error {
	var o orm.Ormer
	o = orm.NewOrm()
	//查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	var querySql = "select " +
		"	sb.bdi_id, " +
		"	sb.bdi_name, " +
		"	sb.bdi_type_id, " +
		"	sb.bdi_secrecy_id, " +
		"	sb.remarks, " +
		"	sbt.bdi_type_name, " +
		"	sbsec.bdi_secrecy_name, " +
		"	group_concat(sbset.bdi_set_id) as bdi_set_ids, " +
		"	group_concat(sbset.bdi_set_name) as bdi_set_names " +
		"from " +
		"	sdt_bdi sb, " +
		"	sdt_bdi_set_rel_bdi rel, " +
		"	sdt_bdi_set sbset, " +
		"	sdt_bdi_type sbt, " +
		"	sdt_bdi_secrecy sbsec " +
		"where " +
		"sb.bdi_id = ? " +
		"and sb.bdi_id = rel.bdi_id " +
		"and rel.bdi_set_id = sbset.bdi_set_id " +
		"and sb.bdi_type_id = sbt.bdi_type_id " +
		"and sb.bdi_secrecy_id = sbsec.bdi_secrecy_id "

	err := o.Raw(querySql, this.BdiId).QueryRow(this)
	if err != nil {
		log.Fatal("查询表：sdt_bdi_set出错！")
		return err
	}

	return nil
}

func (this *SdtBdi) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertSdtBdiSql = " insert into sdt_bdi(bdi_name, bdi_type_id, bdi_secrecy_id, remarks, create_user_id, create_time) values (?, ?, ?, ?, ?, ?)"
	var insertSdtBdiSetRelBdiSql = " insert into sdt_bdi_set_rel_bdi(bdi_set_id, bdi_id) values (?, ?) "

	res, err := o.Raw(insertSdtBdiSql, this.BdiName, this.BdiTypeId, this.BdiSecrecyId, this.Remarks, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	sdtBidId, err := res.LastInsertId()

	if err != nil {
		o.Rollback()
		return err
	}

	bdiSetIds := strings.Split(this.BdiSetIds, ",")
	for _, v := range bdiSetIds {
		bdiSetId, err := strconv.Atoi(v)
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.Raw(insertSdtBdiSetRelBdiSql, bdiSetId, sdtBidId).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	o.Commit()
	return nil
}

func (this *SdtBdi) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var deleteSdtBdiSetRelBdiSql = " delete from sdt_bdi_set_rel_bdi where bdi_id = ? "
	var updateSdtBdiSql = "update sdt_bdi " +
		"set  " +
		" bdi_name = ?, " +
		" bdi_type_id = ?, " +
		" bdi_secrecy_id = ?, " +
		" remarks = ?, " +
		//" modify_user_id = ?, " + //暂时没有修改人
		" modify_time = ? " +
		"where bdi_id = ?"
	var insertSdtBdiSetRelBdiSql = " insert into sdt_bdi_set_rel_bdi(bdi_set_id, bdi_id) values (?, ?) "

	_, err := o.Raw(updateSdtBdiSql, this.BdiName, this.BdiTypeId, this.BdiSecrecyId, this.Remarks, time.Now(), this.BdiId).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	_, err = o.Raw(deleteSdtBdiSetRelBdiSql, this.BdiId).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	bdiSetIds := strings.Split(this.BdiSetIds, ",")

	for _, v := range bdiSetIds {
		bdiSetId, err := strconv.Atoi(v)
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.Raw(insertSdtBdiSetRelBdiSql, bdiSetId, this.BdiId).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	o.Commit()
	return nil
}
