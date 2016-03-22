package controllers

import (
	"bdi/models"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	//"encoding/json"

)

type DatasourceController struct {
	BaseController
}

// 首页
func (this *DatasourceController) Index() {
	this.TplName = "datasource/datasourceIndex.html"
}

/**
根据行数和列数，查询符合条件的指标集的数据。
*/
func (this *DatasourceController) All() {
	var err error
	var rows int
	var page int

	//返回的json数据
	returnData := struct {
		Total interface{} `json:"total"`
		Rows  interface{} `json:"rows"`
	}{}

	//行数
	rows, err = strconv.Atoi(this.GetString("rows"))
	if err != nil {
		rows = 10
	}

	//页数
	page, err = strconv.Atoi(this.GetString("page"))
	if err != nil {
		page = 1
	}

	datasource := new(models.Datasource)
	datasourceSlice, num, err := datasource.GetAllDatasource(rows, page) //查询指标域所有记录

	if err != nil {
		fmt.Println("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = nil
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &datasourceSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 新增Dialog
func (this *DatasourceController) AddPage() {
	this.TplName = "datasource/addDialog.html"
}
//新增
func (this *DatasourceController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	datasource := new(models.Datasource)
	err := this.ParseForm(datasource)
	//err := json.Unmarshal(this.Ctx.Input.RequestBody, datasource)
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = datasource.Add()
	if err != nil {
		returnData.Success = false
		returnData.Message = "新增数据出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Success = true
	returnData.Message = "新增数据成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 更新Dialog
func (this *DatasourceController) UpdatePage() {
	bdiDatasourceId, err := this.GetInt("bdiDatasourceId")
	if err != nil {
		fmt.Println("解析参数出错！")
		return
	}
	datasource := new(models.Datasource)
	datasource.BdiDatasourceId = bdiDatasourceId
	err = datasource.GetDatasourceById()

	if err != nil {
		fmt.Println("解析参数出错！")
		return
	}

	this.Data["datasource"] = datasource
	this.TplName = "datasource/updateDialog.html"
}

//更新信息
func (this *DatasourceController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	datasource := new(models.Datasource)
	err := this.ParseForm(datasource)

	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = datasource.Update()
	if err != nil {
		returnData.Success = false
		returnData.Message = "数据更新出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Success = true
	returnData.Message = "数据更新成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//删除
func (this *DatasourceController) Delete() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	//	bdiDomainId, err := this.GetInt("bdiDomainId")
	//	fmt.Println("bdiDomainId: ", bdiDomainId)
	//	if err != nil {
	//		fmt.Println("解析参数出错")
	//		returnData.Success = false
	//		returnData.Message = "解析参数出错"
	//		this.Data[JSON_STRING] = returnData
	//		this.ServeJSON()
	//		return
	//	}
	//
	//	o := orm.NewOrm()
	//	if _, err := o.Delete(&models.SdtBdiDomain{BdiDomainId: bdiDomainId}); err != nil {
	//		fmt.Println("数据删除失败")
	//		returnData.Success = false
	//		returnData.Message = "数据删除失败"
	//		this.Data[JSON_STRING] = returnData
	//		this.ServeJSON()
	//		return
	//	}

	returnData.Success = true
	returnData.Message = "数据删除成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//检测数据源有效性校验
func (this *DatasourceController) Check() {
	type DataType struct {
		Id   string    `json:"id"`
		Text string `json:"text"`
	}

	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		DatabaseInfo []DataType `json:"databaseInfo"`
	}{}

	ip := this.GetString("ip")
	port := this.GetString("port")
	username := this.GetString("username")
	password := this.GetString("password")
	databaseType := this.GetString("databaseType")

	connectionInfo := username + ":" + password+ "@tcp(" + ip + ":" + port + ")/mysql"

	db, err := sql.Open(databaseType, connectionInfo)
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "数据库连接信息有误！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "数据库连接信息有误！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	rows, err := db.Query(" select s.schema_name as id, s.schema_name as text from information_schema.schemata s ")
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "数据库连接信息有误！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	databaseInfo := new([]DataType)

	for rows.Next() {
		var id string
		var text string

		err := rows.Scan(&id, &text)
		if err != nil {
			fmt.Println(err)
			returnData.Success = false
			returnData.Message = "数据库连接信息有误！"
			this.Data[JSON_STRING] = returnData
			this.ServeJSON()
			return
		}
		dataType := new(DataType)
		dataType.Id = id
		dataType.Text = text
		*databaseInfo = append(*databaseInfo, *dataType)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "数据库连接信息有误！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	rows.Close()

	//在这里进行一些数据库操作
	db.Close()

	returnData.DatabaseInfo = *databaseInfo
	returnData.Success = true
	returnData.Message = "连接成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

func (this *DatasourceController) GetAllDatasourceForTree() {
	datasource := new(models.Datasource)
	datasourceSlice, _ := datasource.GetAllDatasourceForTree()
	this.Data[JSON_STRING] = datasourceSlice
	this.ServeJSON()
	return
}

func (this *DatasourceController) GetAllSchemaForTree() {
	datasource := new(models.Datasource)
	schemaSlice, _ := datasource.GetAllSchemaForTree()
	this.Data[JSON_STRING] = schemaSlice
	this.ServeJSON()
	return
}

func (this *DatasourceController) GetAllTableForTree() {
	schema := this.GetString("param")
	datasource := new(models.Datasource)
	schemaSlice, _ := datasource.GetAllTableForTree(schema)
	this.Data[JSON_STRING] = schemaSlice
	this.ServeJSON()
	return
}


func (this *DatasourceController) GetAllColumnForTree() {
	table := this.GetString("param")

	datasource := new(models.Datasource)
	schemaSlice, _ := datasource.GetAllColumnForTree(table)
	this.Data[JSON_STRING] = schemaSlice
	this.ServeJSON()
	return
}