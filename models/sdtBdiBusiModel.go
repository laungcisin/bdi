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

type BusiTreeAttributes struct {
	BusiId       int                 `form:"busiId"`
	Params       string                 `form:"params"`
	ProcessType  string                 `form:"processType"`
	Name         string                 `form:"name"`
	CnName       string                 `form:"cnName"`
	BdiId        int                    `form:"bdiId"`
	DatasourceId int                    `form:"datasourceId"`
	ChildColumns []ColumnTreeAttributes `form:"childColumns"`
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

func (this *SdtBdiBusi) Update(busiTreeAttributes []BusiTreeAttributes) error {
	var err error
	var latestBusiId int64
	o := orm.NewOrm()
	o.Begin()

	var updateSql = "update sdt_bdi_busi set  " +
		" process_type = ?, " +
		" user_code = ?, " +
		" edit_time = ? " +
		" where id = ?"

	var insertBusiSql = " insert into sdt_bdi_busi(bdi_id, datasource_id, name, cn_name, create_time) values (?, ?, ?, ?, ?)"
	var insertBusiConfigSql =  " insert into sdt_bdi_busi_config(busi_id, cn_name, process_column, process_data_type, process_data_length, user_code, create_time) " +
	" values(?, ?, ?, ?, ?, ?, ?) "

	for _, tableValue := range busiTreeAttributes {
		_, err = o.Raw(updateSql, tableValue.ProcessType, 0, time.Now(), tableValue.BusiId).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		var sdtBdiBusiSlice []SdtBdiBusi = make([]SdtBdiBusi, 0)
		var querySql = " select * from sdt_bdi_busi b where b.bdi_id = ? and b.name = ? "
		_, err = o.Raw(querySql, tableValue.BdiId, tableValue.Name).QueryRows(&sdtBdiBusiSlice)
		if err == nil {
			if len(sdtBdiBusiSlice) > 0 { //有记录
				latestBusiId = int64(tableValue.BusiId)
			}else {
				tableResult, err := o.Raw(insertBusiSql, tableValue.BdiId, tableValue.DatasourceId, tableValue.Name, tableValue.CnName, time.Now()).Exec()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}

				latestBusiId, err = tableResult.LastInsertId()
				if err != nil {
					fmt.Println(err)
					o.Rollback()
					return err
				}

			}
		} else {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		for _, fieldValue := range tableValue.ChildColumns {
			var querySql = " select count(*) as counts from sdt_bdi_busi_config config where config.busi_id = ? and config.process_column = lower(?) "
			num := new(int)
			err := o.Raw(querySql, latestBusiId, tableValue.Name).QueryRow(&num)
			if err == nil {
				if *num > 0 {
					continue
				}
			} else {
				fmt.Println(err)
				o.Rollback()
				return err
			}

			_, err = o.Raw(insertBusiConfigSql, latestBusiId, fieldValue.Comment, fieldValue.Name, fieldValue.DataType,
				fieldValue.DataLength, 0, time.Now()).Exec()

			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}
	}

	o.Commit()
	return nil
}

func (this *SdtBdiBusi) AddBusiAndAddBusiConfig(tableTreeAttributes []TableTreeAttributes) error {
	o := orm.NewOrm()
	o.Begin()

	var insertBusiSql = " insert into sdt_bdi_busi(bdi_id, datasource_id, name, cn_name, create_time) values (?, ?, ?, ?, ?)"

	for _, tableValue := range tableTreeAttributes {
		var querySql = " select count(*) as counts from sdt_bdi_busi b where b.bdi_id = ? and b.name = ? "
		num := new(int)
		err := o.Raw(querySql, tableValue.BdiId, tableValue.Name).QueryRow(&num)
		if err == nil {
			if *num > 0 {
				continue
			}
		} else {
			fmt.Println(err)
			o.Rollback()
			return err
		}

		_, err = o.Raw(insertBusiSql, tableValue.BdiId, tableValue.DatasourceId, tableValue.Name, tableValue.CnName, time.Now()).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}

	o.Commit()
	return nil
}

// sdt_bdi_busi - 详细配置
func (this *SdtBdiBusi) UpdateDetailConfig() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql = "update sdt_bdi_busi " +
	" set  " +
	" is_process = ?, " +
	" process_type = ?, " +
	" params = ?, " +
	" user_code = ?, " +
	" edit_time = ? " +
	" where id = ?"

	_, err := o.Raw(updateSql, this.IsProcess, this.ProcessType, this.Params, '0', time.Now(), this.Id).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}