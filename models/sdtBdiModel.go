package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"strings"
	"time"
)

type SdtBdi struct {
	Id        int    `form:"bdiId"`     //主键
	BdiTypeId int    `form:"bdiTypeId"` //指标类型ID
	BdiName   string `form:"bdiName"`   //指标名称
	Remarks   string `form:"remarks"`   //指标集备注

	UserCode   string    `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//新增时使用
	BdiSetIds string `form:"bdiSetIds"` //所属指标集ids

	//以下字段为datagrid展示
	TypeName    string //所属指标类型名称
	BdiSetNames string //所属指标集名称
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
	var sdtBdiSlice []SdtBdi = make([]SdtBdi, 0)

	//查询的字段顺序最好和model的字段对应，方便解析并赋值。
	var querySql = "select " +
		"	sb.id, " +
		"	sb.bdi_name, " +
		"	sb.bdi_type_id, " +
		"	sb.remarks, " +
		"	sbt.type_name, " +
		"	group_concat(sbset.id) as bdi_set_ids, " +
		"	group_concat(sbset.bdi_set_name) as bdi_set_names " +
		"from " +
		"	sdt_bdi sb, " +
		"	sdt_bdi_set_bdi rel, " +
		"	sdt_bdi_set sbset, " +
		"	sdt_bdi_type sbt " +
		"where " +
		"	sb.id = rel.bdi_id " +
		"and rel.set_id = sbset.id " +
		"and sb.bdi_type_id = sbt.id " +
		"group by sb.id limit ?, ? "

	num, err := o.Raw(querySql, (page-1)*rows, page*rows).QueryRows(&sdtBdiSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	var countSql = "select " +
		"	count(*) as counts " +
		"from " +
		"	sdt_bdi sb, " +
		"	sdt_bdi_set_bdi rel, " +
		"	sdt_bdi_set sbset, " +
		"	sdt_bdi_type sbt " +
		"where " +
		"	sb.id = rel.bdi_id " +
		"and rel.set_id = sbset.id " +
		"and sb.bdi_type_id = sbt.id"

	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
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
		"	sb.id, " +
		"	sb.bdi_name, " +
		"	sb.bdi_type_id, " +
		"	sb.remarks, " +
		"	sbt.type_name, " +
		"	group_concat(sbset.id) as bdi_set_ids, " +
		"	group_concat(sbset.bdi_set_name) as bdi_set_names " +
		"from " +
		"	sdt_bdi sb, " +
		"	sdt_bdi_set_bdi rel, " +
		"	sdt_bdi_set sbset, " +
		"	sdt_bdi_type sbt " +
		"where " +
		" sb.id = ? " +
		" and sb.id = rel.bdi_id " +
		" and rel.set_id = sbset.id " +
		" and sb.bdi_type_id = sbt.id "

	err := o.Raw(querySql, this.Id).QueryRow(this)
	if err != nil {
		fmt.Println(err)
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return err
	}

	return nil
}

func (this *SdtBdi) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertSdtBdiSql = " insert into sdt_bdi(bdi_name, bdi_type_id, remarks, user_code, create_time) values (?, ?, ?, ?, ?)"
	var insertSdtBdiSetRelBdiSql = " insert into sdt_bdi_set_bdi(set_id, bdi_id) values (?, ?) "

	res, err := o.Raw(insertSdtBdiSql, this.BdiName, this.BdiTypeId, this.Remarks, 0, time.Now()).Exec()
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

	fmt.Println("SdtBdi: ", this)
	var deleteSdtBdiSetRelBdiSql = " delete from sdt_bdi_set_bdi where bdi_id = ? "
	var updateSdtBdiSql = "update sdt_bdi " +
		"set  " +
		" bdi_name = ?, " +
		" bdi_type_id = ?, " +
		" remarks = ?, " +
		" edit_time = ? " +
		"where id = ?"
	var insertSdtBdiSetRelBdiSql = " insert into sdt_bdi_set_bdi(set_id, bdi_id) values (?, ?) "

	_, err := o.Raw(updateSdtBdiSql, this.BdiName, this.BdiTypeId, this.Remarks, time.Now(), this.Id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	_, err = o.Raw(deleteSdtBdiSetRelBdiSql, this.Id).Exec()
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
		_, err = o.Raw(insertSdtBdiSetRelBdiSql, bdiSetId, this.Id).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	o.Commit()
	return nil
}
