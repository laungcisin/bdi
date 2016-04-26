package models

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	"log"
	//"strconv"
	//"strings"
	"time"
	//"encoding/json"
	"strconv"
	"strings"
	//"fmt"
	"encoding/json"
)

const (
	UrlForDatasource = "/datasource/datasourceNode"
	UrlForSchema     = "/datasource/schemaNode"
	UrlForTable      = "/datasource/tableNode"
	UrlForColumn     = "/datasource/columnNode"
)

type Datasource struct {
	Id       int    `form:"id"`       //主键
	TypeId   int    `form:"typeId"`   //数据源类型Id
	Name     string `form:"name"`     //数据源名称
	Ip       string `form:"ip"`       //ip地址
	Port     int    `form:"port"`     //端口号
	Dbname   string `form:"dbname"`   //数据库名
	Username string `form:"username"` //用户名
	Password string `form:"password"` //密码
	Driver   string `form:"driver"`   //数据库驱动
	Remarks  string `form:"remarks"`  //数据源备注

	UserCode   int       `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//以下字段为datagrid展示
	TypeName string //数据源类型名称
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
	var datasourceSlice []Datasource = make([]Datasource, 0)

	//查询的字段顺序最好和model的字段对应，方便解析并赋值。
	var querySql = " select d.*, t.type_name " +
		" from  sdt_bdi_datasource d " +
		" left join sdt_bdi_datasource_type t on d.type_id = t.id " +
		" limit ?, ? "

	_, err := o.Raw(querySql, (page-1)*rows, page*rows).QueryRows(&datasourceSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	num := new(int)
	var countSql = " select count(*) as counts  " +
		" from  sdt_bdi_datasource d " +
		" left join sdt_bdi_datasource_type t on d.type_id = t.id "

	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return datasourceSlice, *num, nil
}

/**
获取所有指标。
*/
func (this *Datasource) GetDatasourceById() error {
	o := orm.NewOrm()
	//查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	var querySql = " select d.*, t.type_name " +
		" from sdt_bdi_datasource d " +
		" left join sdt_bdi_datasource_type t on d.type_id = t.id " +
		" where d.id = ? "

	err := o.Raw(querySql, this.Id).QueryRow(this)

	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return err
	}

	return nil
}

func (this *Datasource) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertSdtBdiSql = " insert into sdt_bdi_datasource(type_id, name, ip, port, dbname, username, password, driver, remarks, user_code, create_time)values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	_, err := o.Raw(insertSdtBdiSql, this.TypeId, this.Name, this.Ip, this.Port, this.Dbname, this.Username, this.Password,
		this.Driver, this.Remarks, 0, time.Now()).Exec()
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
		" update sdt_bdi_datasource " +
		" set type_id = ?, " +
		" name = ?, " +
		" ip = ?, " +
		" port = ?, " +
		" dbname = ?, " +
		" username = ?, " +
		" password = ?, " +
		" driver = ?, " +
		" remarks = ?, " +
		//" user_code = '0', " +
		" edit_time = ? " +
		" where id = ? "
	_, err := o.Raw(updateSdtBdiSql, this.TypeId, this.Name, this.Ip, this.Port, this.Dbname, this.Username, this.Password,
		this.Driver, this.Remarks, time.Now(), this.Id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

type DatasourceTreeAttributes struct {
	Url        string `json:"url"`
	Ip         string `json:"ip"`
	Port       string `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	TableName  string `json:"tableName"`
	SchemaName string `json:"schemaName"`
}

type DatasourceTree struct {
	Id      string `json:"id"`
	Pid     string `json:"pid"`
	Text    string `json:"text"`
	IconCls string `json:"iconCls"`
	Checked string `json:"checked"`
	State   string `json:"state"`

	Attributes DatasourceTreeAttributes `json:"attributes"`
	Children   []Tree                   `json:"children"`
}

/**
获取所有数据源--用于tree。
*/
func (this *Datasource) GetAllDatasourceForTree() ([]DatasourceTree, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	sdtBdiDatasourceSlice := make([]Datasource, 0, 10)

	var querySql = " select * from " + this.TableName()

	_, err := o.Raw(querySql).QueryRows(&sdtBdiDatasourceSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, err
	}

	newTreeDataSlice := make([]DatasourceTree, 0)
	for _, v := range sdtBdiDatasourceSlice {
		treeNode := new(DatasourceTree)
		treeNode.Id = strconv.Itoa(v.Id)
		treeNode.Text = v.Name
		treeNode.State = "closed"

		var connectionInfo string
		// connectionInfo = strings.Replace(v.DatasourceConnect, "#0F01", "{", -1)
		connectionInfo = strings.Replace(connectionInfo, "#0F02", "\"", -1)
		connectionInfo = strings.Replace(connectionInfo, "#0F03", "}", -1)
		datasourceTreeAttributes := new(DatasourceTreeAttributes)
		json.Unmarshal([]byte(connectionInfo), datasourceTreeAttributes)
		datasourceTreeAttributes.Url = UrlForSchema
		treeNode.Attributes = *datasourceTreeAttributes
		newTreeDataSlice = append(newTreeDataSlice, *treeNode)
	}

	return newTreeDataSlice, nil
}

func (this *Datasource) GetAllSchemaForTree(ip string, port string, username string, password string, tableName string, schemaName string) ([]DatasourceTree, error) {
	newTreeDataSlice := make([]DatasourceTree, 0, 10)
	connectionInfo := username + ":" + password + "@tcp(" + ip + ":" + port + ")/mysql"

	db, err := sql.Open("mysql", connectionInfo)
	if err != nil {
		return newTreeDataSlice, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return newTreeDataSlice, err
	}

	rows, err := db.Query(" select s.schema_name as id, s.schema_name as text from information_schema.schemata s ")
	if err != nil {
		return newTreeDataSlice, err
	}

	for rows.Next() {
		var id string
		var text string
		err := rows.Scan(&id, &text)
		if err != nil {
			return newTreeDataSlice, err
		}
		treeNode := new(DatasourceTree)
		treeNode.Id = id
		treeNode.Text = text
		treeNode.State = "closed"
		datasourceTreeAttributes := new(DatasourceTreeAttributes)
		datasourceTreeAttributes.Ip = ip
		datasourceTreeAttributes.Port = port
		datasourceTreeAttributes.Username = username
		datasourceTreeAttributes.Password = password
		datasourceTreeAttributes.Url = UrlForTable
		treeNode.Attributes = *datasourceTreeAttributes
		newTreeDataSlice = append(newTreeDataSlice, *treeNode)
	}
	defer rows.Close()
	return newTreeDataSlice, nil
}

func (this *Datasource) GetAllTableForTree(ip string, port string, username string, password string, tableName string, schemaName string) ([]DatasourceTree, error) {
	newTreeDataSlice := make([]DatasourceTree, 0, 10)

	connectionInfo := username + ":" + password + "@tcp(" + ip + ":" + port + ")/mysql"
	db, err := sql.Open("mysql", connectionInfo)
	if err != nil {
		return newTreeDataSlice, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return newTreeDataSlice, err
	}

	rows, err := db.Query(" select table_name as id, table_name as text from information_schema.tables where table_schema = ? ", schemaName)
	if err != nil {
		return newTreeDataSlice, err
	}

	for rows.Next() {
		var id string
		var text string
		err := rows.Scan(&id, &text)
		if err != nil {
			return newTreeDataSlice, err
		}
		treeNode := new(DatasourceTree)
		treeNode.Id = id
		treeNode.Text = text
		treeNode.State = "closed"
		datasourceTreeAttributes := new(DatasourceTreeAttributes)
		datasourceTreeAttributes.Ip = ip
		datasourceTreeAttributes.Port = port
		datasourceTreeAttributes.Username = username
		datasourceTreeAttributes.Password = password
		datasourceTreeAttributes.Url = UrlForColumn
		datasourceTreeAttributes.SchemaName = schemaName
		treeNode.Attributes = *datasourceTreeAttributes
		newTreeDataSlice = append(newTreeDataSlice, *treeNode)
	}
	defer rows.Close()
	return newTreeDataSlice, nil
}

func (this *Datasource) GetAllColumnForTree(ip string, port string, username string, password string, tableName string, schemaName string) ([]DatasourceTree, error) {
	newTreeDataSlice := make([]DatasourceTree, 0, 10)

	connectionInfo := username + ":" + password + "@tcp(" + ip + ":" + port + ")/mysql"

	db, err := sql.Open("mysql", connectionInfo)
	if err != nil {
		return newTreeDataSlice, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return newTreeDataSlice, err
	}

	rows, err := db.Query(" select c.column_name as id , c.column_name as text from information_schema.columns c where c.table_schema = ? and c.table_name = ? ", schemaName, tableName)
	if err != nil {
		return newTreeDataSlice, err
	}

	for rows.Next() {
		var id string
		var text string
		err := rows.Scan(&id, &text)
		if err != nil {
			return newTreeDataSlice, err
		}
		treeNode := new(DatasourceTree)
		treeNode.Id = id
		treeNode.Text = text
		treeNode.State = "open"
		datasourceTreeAttributes := new(DatasourceTreeAttributes)
		datasourceTreeAttributes.Url = ""
		datasourceTreeAttributes.Ip = ip
		datasourceTreeAttributes.Port = port
		datasourceTreeAttributes.Username = username
		datasourceTreeAttributes.Password = password
		datasourceTreeAttributes.SchemaName = schemaName
		datasourceTreeAttributes.TableName = tableName
		treeNode.Attributes = *datasourceTreeAttributes
		newTreeDataSlice = append(newTreeDataSlice, *treeNode)
	}
	defer rows.Close()
	return newTreeDataSlice, nil
}
