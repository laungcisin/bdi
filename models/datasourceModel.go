package models

import (
	"github.com/astaxie/beego/orm"
	"log"
//"strconv"
//"strings"
	"time"
	"fmt"
	"encoding/json"
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

	//以下字段接收后，以json格式存储在DatasourceConnect字段中
	Ip string `form:"ip"`             //ip地址
	Port string `form:"port"`             //端口号
	Username string `form:"username"`             //数据库用户名
	Password string `form:"password"`             //数据库密码
	Schemata string `form:"schemata"`             //数据库库名
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
	type Config struct {
		Ip string `form:"ip"`             //ip地址
		Port string `form:"port"`             //端口号
		Username string `form:"username"`             //数据库用户名
		Password string `form:"password"`             //数据库密码
		Schemata string `form:"schemata"`             //数据库库名
	}

	var config Config = Config{}
	config.Ip = this.Ip
	config.Port = this.Port
	config.Username = this.Username
	config.Password = this.Password
	config.Schemata = this.Schemata

	result, err := json.Marshal(config)
	if err != nil {
		return err
	}

	fmt.Println(string(result))
	this.DatasourceConnect = string(result)

	o := orm.NewOrm()
	o.Begin()

	var insertSdtBdiSql = " insert into sdt_bdi_datasource(bdi_datasource_name, bdi_datasource_type_id, datasource_connect, remarks, create_user_id, create_time)values(?, ?, ?, ?, ?, ?) "

	_, err = o.Raw(insertSdtBdiSql, this.BdiDatasourceName, this.BdiDatasourceTypeId, this.DatasourceConnect, this.Remarks, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

func (this *Datasource) Update() error {
	//o := orm.NewOrm()
	//o.Begin()
	//
	//var deleteSdtBdiSetRelBdiSql = " delete from sdt_bdi_set_rel_bdi where bdi_id = ? "
	//var updateSdtBdiSql = "update sdt_bdi " +
	//	"set  " +
	//	" bdi_name = ?, " +
	//	" bdi_type_id = ?, " +
	//	" bdi_secrecy_id = ?, " +
	//	" remarks = ?, " +
	//	//" modify_user_id = ?, " + //暂时没有修改人
	//	" modify_time = ? " +
	//	"where bdi_id = ?"
	//var insertSdtBdiSetRelBdiSql = " insert into sdt_bdi_set_rel_bdi(bdi_set_id, bdi_id) values (?, ?) "
	//
	//_, err := o.Raw(updateSdtBdiSql, this.BdiName, this.BdiTypeId, this.BdiSecrecyId, this.Remarks, time.Now(), this.BdiId).Exec()
	//if err != nil {
	//	o.Rollback()
	//	return err
	//}
	//
	//_, err = o.Raw(deleteSdtBdiSetRelBdiSql, this.BdiId).Exec()
	//if err != nil {
	//	o.Rollback()
	//	return err
	//}
	//
	//bdiSetIds := strings.Split(this.BdiSetIds, ",")
	//
	//for _, v := range bdiSetIds {
	//	bdiSetId, err := strconv.Atoi(v)
	//	if err != nil {
	//		o.Rollback()
	//		return err
	//	}
	//	_, err = o.Raw(insertSdtBdiSetRelBdiSql, bdiSetId, this.BdiId).Exec()
	//	if err != nil {
	//		o.Rollback()
	//		return err
	//	}
	//}
	//
	//o.Commit()
	return nil
}
