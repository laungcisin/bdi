package models

import (
	"github.com/astaxie/beego/orm"
	"log"
//"strconv"
//"strings"
	"time"
	//"encoding/json"
)

type Datasource struct {
	BdiDatasourceId       int    `form:"bdiDatasourceId"`     //主键
	BdiDatasourceName     string `form:"bdiDatasourceName"`   //数据源名称
	BdiDatasourceTypeId   int    `form:"bdiDatasourceTypeId"` //数据源类型Id
	DatasourceConnect     string    `form:"datasourceConnect"`   //数据源连接详情
	Remarks               string `form:"remarks"`             //数据源备注

	CreateUserId          int                                 //创建人ID
	CreateTime            time.Time                           //创建时间
	ModifyUserId          int                                 //修改人ID
	ModifyTime            time.Time                           //修改时间

								  //以下字段为datagrid展示
	BdiDatasourceTypeName string                              //数据源类型名称
}

func (this *Datasource) TableName() string {
	return "sdt_bdi_datasource"
}

/**
获取所有数据源。
*/
func (this *Datasource) GetAllDatasource(rows int, page int) ([]Datasource, int, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	var datasourceSlice []Datasource

	//查询的字段顺序最好和model的字段对应，方便解析并赋值。
	var querySql =
	" select d.*, t.bdi_datasource_type_name " +
	" from  sdt_bdi_datasource d " +
	" left join sdt_bdi_datasource_type t on d.bdi_datasource_type_id = t.bdi_datasource_type_id " +
	" where d.bdi_datasource_id limit ?, ? "

	num, err := o.Raw(querySql, (page - 1) * rows, page * rows).QueryRows(&datasourceSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	var countSql = " select count(*) as counts from " + this.TableName()
	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return datasourceSlice, int(num), nil
}

/**
获取所有指标。
*/
func (this *Datasource) GetDatasourceById() error {
	var o orm.Ormer
	o = orm.NewOrm()
	//查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	var querySql =
	" select d.*, t.bdi_datasource_type_name " +
	" from sdt_bdi_datasource d " +
	" left join sdt_bdi_datasource_type t on d.bdi_datasource_type_id = t.bdi_datasource_type_id " +
	" where d.bdi_datasource_id = ? "

	err := o.Raw(querySql, this.BdiDatasourceId).QueryRow(this)

	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return err
	}

	return nil
}

func (this *Datasource) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertSdtBdiSql = " insert into sdt_bdi_datasource(bdi_datasource_name, bdi_datasource_type_id, datasource_connect, remarks, create_user_id, create_time)values(?, ?, ?, ?, ?, ?) "

	_, err := o.Raw(insertSdtBdiSql, this.BdiDatasourceName, this.BdiDatasourceTypeId, this.DatasourceConnect, this.Remarks, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

func (this *Datasource) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSdtBdiSql =
		"update sdt_bdi_datasource " +
		"set  " +
		" bdi_datasource_name = ?, " +
		" bdi_datasource_type_id = ?, " +
		" datasource_connect = ?, " +
		" remarks = ?, " +
		//" modify_user_id = ?, " +//暂时没有修改人
		" modify_time = ? " +
		"where bdi_datasource_id = ?"

	_, err := o.Raw(updateSdtBdiSql, this.BdiDatasourceName, this.BdiDatasourceTypeId, this.DatasourceConnect, this.Remarks, time.Now(), this.BdiDatasourceId).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}
