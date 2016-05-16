package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type SdtBdiBusiConfig struct {
	Id                int    `form:"id"`                //主键
	BusiId            int    `form:"busiId"`            //业务表ID
	Name              string `form:"name"`              //配置业务表名
	CnName            string `form:"cnName"`            //配置业务表中文名
	IsProcess         int    `form:"isProcess"`         //是否经过处理
	SelectStar        string `form:"selectStar"`        //选择字段
	ProcessColumn     string `form:"processColumn"`     //处理字段
	AsName            string `form:"asName"`            //别名
	ProcessDataType   string `form:"processDataType"`   //
	ProcessDataLength string `form:"processDataLength"` //
	ProcessType       string `form:"processType"`       //处理类型
	Params            string `form:"params"`            //参数

	UserCode   string    `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//以下字段为datagrid展示
	BdiId int `form:"bdiId"` //创建人ID

}

func (u *SdtBdiBusiConfig) TableName() string {
	return "sdt_bdi_busi_config"
}

/**
获取所有指标。
*/
func (this *SdtBdiBusiConfig) GetAllSdtBdiBusiConfig(rows int, page int) ([]SdtBdiBusiConfig, int, error) {
	o := orm.NewOrm()
	var SdtBdiBusiSlice []SdtBdiBusiConfig = make([]SdtBdiBusiConfig, 0)

	var querySql = " select " +
		"	busi.*, " +
		"   	bdi.bdi_name as bdi_name, " +
		"	d.name as datasource_name " +
		" from " +
		"	sdt_bdi_busi_cfg busi, " +
		"	sdt_bdi bdi, " +
		"	sdt_bdi_datasource d " +
		" where " +
		" 	busi.bdi_id = ? " +
		"	and busi.bdi_id = bdi.id " +
		" 	and busi.datasource_id = d.id " +
		"  limit ?, ? "

	_, err := o.Raw(querySql, this.BdiId, (page-1)*rows, page*rows).QueryRows(&SdtBdiBusiSlice)
	if err != nil {
		fmt.Println("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	num := new(int)
	var countSql = "select " +
		"	count(*) as counts " +
		" from " +
		"	sdt_bdi_busi_cfg busi, " +
		"	sdt_bdi bdi, " +
		"	sdt_bdi_datasource d " +
		" where " +
		" 	busi.bdi_id = ? " +
		"	and busi.bdi_id = bdi.id " +
		" 	and busi.datasource_id = d.id "

	err = o.Raw(countSql, this.BdiId).QueryRow(&num)
	if err != nil {
		fmt.Println("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return SdtBdiBusiSlice, *num, nil
}

/**
获取所有指标。
*/
//func (this *SdtBdiBusiConfig) GetSdtBdiBusiConfigById() error {
//	o := orm.NewOrm()
//	var querySql = " select busi.* from sdt_bdi_busi_cfg busi where busi.id = ? "
//
//	err := o.Raw(querySql, this.Id).QueryRow(this)
//	if err != nil {
//		log.Fatal("查询表：" + this.TableName() + "出错！")
//		return err
//	}
//
//	fmt.Println(this)
//	return nil
//}

/*
func (this *SdtBdiBusiConfig) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertSdtBdiBusiSql = "insert into sdt_bdi_busi_cfg ( " +
		"	bdi_id, " +
		"	datasource_id, " +
		"	name, " +
		"	user_code, " +
		"	create_time " +
		") " +
		"values (?, ?, ?, ?, ?) "

	_, err := o.Raw(insertSdtBdiBusiSql, this.BdiId, this.BdiId, this.Name, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (this *SdtBdiBusiConfig) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql = "update sdt_bdi_busi_cfg " +
		" set  " +
		" datasource_id = ?, " +
		" name = ?, " +
		" user_code = ?, " +
		" edit_time = ? " +
		" where id = ?"

	_, err := o.Raw(updateSql, this.BdiId, this.Name, this.UserCode, time.Now(), this.Id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (this *SdtBdiBusiConfig) AddBusiAndAddField(tableTreeAttributes []TableTreeAttributes) error {
	o := orm.NewOrm()
	o.Begin()

	//num := new(int)
	//var countSql = " select count(*) as counts from sdt_bdi_busi_cfg b where b.bdi_id = ? and b.name = ? "
	//err := o.Raw(countSql).QueryRow(num)

	var insertTableSql = " insert into sdt_bdi_busi_cfg(bdi_id, datasource_id, name, cn_name, create_time) values (?, ?, ?, ?, ?)"
	var insertFieldSql = " insert into sdt_bdi_busi_fields(busi_id, name, sequence, comment, data_type, data_length, user_code, create_time) values (?, ?, ?, ?, ?, ?, ?, ?)"

	var sequenceIndex int = 1

	for _, tableValue := range tableTreeAttributes {
		//先看是否有重复记录，如果有，先删除记录
		var sdtBdiBusiSlice []SdtBdiBusiConfig = make([]SdtBdiBusiConfig, 0)
		var querySql = " select * from sdt_bdi_busi_cfg b where b.bdi_id = ? and b.name = ? "
		_, err := o.Raw(querySql, tableValue.BdiId, tableValue.Name).QueryRows(&sdtBdiBusiSlice)
		if err == nil {
			if len(sdtBdiBusiSlice) > 0 { //如果有记录，先删除记录
				for _, tempTableValue := range sdtBdiBusiSlice {
					_, err = o.Raw(" delete from sdt_bdi_busi_cfg where bdi_id = ? and name = ? ", tempTableValue.BdiId, tempTableValue.Name).Exec()
					if err != nil {
						fmt.Println(err)
						o.Rollback()
						return err
					}

					_, err = o.Raw(" delete from sdt_bdi_busi_fields where busi_id = ? ", tempTableValue.Id).Exec()
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

		tableResult, err := o.Raw(insertTableSql, tableValue.BdiId, tableValue.BdiId, tableValue.Name, tableValue.CnName, time.Now()).Exec()
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
			_, err = o.Raw(insertFieldSql, lastInsertId, fieldValue.Name, sequenceIndex, fieldValue.Comment, fieldValue.DataType,
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
*/
