package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	// "strconv"
	// "strings"
	"fmt"
	"time"
)

type SdtBdiBusi struct {
	Id           int    `form:"id"`           //主键
	BdiId        int    `form:"bdiId"`        //指标id
	DatasourceId int    `form:"datasourceId"` //数据源id
	Name         string `form:"name"`         //业务表名
	CnName       string `form:"CnName"`       //业务中文名
	IsProcess    int    `form:"isProcess"`    //是否经过处理
	ProcessType  string `form:"processType"`  //处理类型
	Params       string `form:"params"`       //参数

	UserCode   string    `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//以下字段为datagrid展示
	BdiName        string
	DatasourceName string
}

func (u *SdtBdiBusi) TableName() string {
	return "sdt_bdi_busi"
}

/**
获取所有指标。
*/
func (this *SdtBdiBusi) GetAllSdtBdiBusi(rows int, page int) ([]SdtBdiBusi, int, error) {
	o := orm.NewOrm()
	var SdtBdiBusiSlice []SdtBdiBusi = make([]SdtBdiBusi, 0)

	var querySql = " select " +
		"	busi.*, " +
		"   	bdi.bdi_name as bdi_name, " +
		"	d.name as datasource_name " +
		" from " +
		"	sdt_bdi_busi busi, " +
		"	sdt_bdi bdi, " +
		"	sdt_bdi_datasource d " +
		" where " +
		" 	busi.bdi_id = ? " +
		"	and busi.bdi_id = bdi.id " +
		" 	and busi.datasource_id = d.id " +
		"  limit ?, ? "

	_, err := o.Raw(querySql, this.BdiId, (page-1)*rows, page*rows).QueryRows(&SdtBdiBusiSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	num := new(int)
	var countSql = "select " +
		"	count(*) as counts " +
		" from " +
		"	sdt_bdi_busi busi, " +
		"	sdt_bdi bdi, " +
		"	sdt_bdi_datasource d " +
		" where " +
		" 	busi.bdi_id = ? " +
		"	and busi.bdi_id = bdi.id " +
		" 	and busi.datasource_id = d.id "

	err = o.Raw(countSql, this.BdiId).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return SdtBdiBusiSlice, *num, nil
}

/**
获取所有指标。
*/
func (this *SdtBdiBusi) GetSdtBdiBusiById() error {
	o := orm.NewOrm()
	var querySql = " select busi.* from sdt_bdi_busi busi where busi.id = ? "

	err := o.Raw(querySql, this.Id).QueryRow(this)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return err
	}

	fmt.Println(this)
	return nil
}

func (this *SdtBdiBusi) Add() error {
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

	_, err := o.Raw(insertSdtBdiBusiSql, this.BdiId, this.DatasourceId, this.Name, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (this *SdtBdiBusi) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql = "update sdt_bdi_busi " +
		" set  " +
		" datasource_id = ?, " +
		" name = ?, " +
		" user_code = ?, " +
		" edit_time = ? " +
		" where id = ?"

	_, err := o.Raw(updateSql, this.DatasourceId, this.Name, this.UserCode, time.Now(), this.Id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (this *SdtBdiBusi) AddBusiAndAddBusiConfig(tableTreeAttributes []TableTreeAttributes) error {
	o := orm.NewOrm()
	o.Begin()

	var insertBusiSql = " insert into sdt_bdi_busi(bdi_id, datasource_id, name, cn_name, create_time) values (?, ?, ?, ?, ?)"
	var insertBusiConfigSql =  " insert into sdt_bdi_busi_config(busi_id, cn_name, process_column, process_data_type, process_data_length, user_code, create_time) " +
	" values(?, ?, ?, ?, ?, ?, ?) "

	var sequenceIndex int = 1
	for _, tableValue := range tableTreeAttributes {
		//先看是否有重复记录，如果有，先删除记录
		var sdtBdiBusiSlice []SdtBdiBusi = make([]SdtBdiBusi, 0)
		var querySql = " select * from sdt_bdi_busi b where b.bdi_id = ? and b.name = ? "
		_, err := o.Raw(querySql, tableValue.BdiId, tableValue.Name).QueryRows(&sdtBdiBusiSlice)
		if err == nil {
			if len(sdtBdiBusiSlice) > 0 { //如果有记录，先删除记录
				for _, tempTableValue := range sdtBdiBusiSlice {
					_, err = o.Raw(" delete from sdt_bdi_busi where bdi_id = ? and name = ? ", tempTableValue.BdiId, tempTableValue.Name).Exec()
					if err != nil {
						fmt.Println(err)
						o.Rollback()
						return err
					}

					_, err = o.Raw(" delete from sdt_bdi_busi_config where busi_id = ? ", tempTableValue.Id).Exec()
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

		tableResult, err := o.Raw(insertBusiSql, tableValue.BdiId, tableValue.DatasourceId, tableValue.Name, tableValue.CnName, time.Now()).Exec()
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
			_, err = o.Raw(insertBusiConfigSql, lastInsertId, fieldValue.Comment, fieldValue.Name, fieldValue.DataType,
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
