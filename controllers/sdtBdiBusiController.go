package controllers

import (
	"github.com/astaxie/beego/orm"
	"bdi/models"
	"fmt"
	"log"
	"strconv"
	"encoding/json"
)

type SdtBdiBusiController struct {
	BaseController
}

// 首页
func (this *SdtBdiBusiController) Index() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	this.TplName = "sdtBdiBusi/sdtBdiBusiIndex.html"
}

/**
根据行数和列数，查询符合条件的数据。
*/
func (this *SdtBdiBusiController) All() {
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

	bdiId, _ := this.GetInt("bdiId")
	sdtBdiBusi := new(models.SdtBdiBusi)
	sdtBdiBusi.BdiId = bdiId
	sdtBdiBusiSlice, num, err := sdtBdiBusi.GetAllSdtBdiBusi(rows, page)
	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = &sdtBdiBusiSlice
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &sdtBdiBusiSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 新增Dialog
func (this *SdtBdiBusiController) AddPage() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	this.TplName = "sdtBdiBusi/addDialog.html"
}

//新增
func (this *SdtBdiBusiController) Add() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var ob []models.TableTreeAttributes
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	sdtBdiBusi := new(models.SdtBdiBusi)
	err = sdtBdiBusi.AddBusiAndAddBusiConfig(ob)
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "保存数据出错"
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
func (this *SdtBdiBusiController) UpdatePage() {
	id, err := this.GetInt("id")
	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}
	sdtBdiBusi := new(models.SdtBdiBusi)
	sdtBdiBusi.Id = id
	err = sdtBdiBusi.GetSdtBdiBusiById()

	if err != nil {
		fmt.Println(err)
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdiBusi"] = sdtBdiBusi
	this.TplName = "sdtBdiBusi/updateDialog.html"
}

//更新信息
func (this *SdtBdiBusiController) Update() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var ob []models.BusiTreeAttributes
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	sdtBdiBusi := new(models.SdtBdiBusi)
	err = sdtBdiBusi.Update(ob)
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
func (this *SdtBdiBusiController) Delete() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	returnData.Success = true
	returnData.Message = "数据删除成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// SdtBdiBusi-详细配置页面。
func (this *SdtBdiBusiController) DetailConfigPage() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	id, _ := this.GetInt("id")
	this.Data["id"] = id

	sdtBdiBusi := new(models.SdtBdiBusi)
	sdtBdiBusi.Id = id
	sdtBdiBusi.GetSdtBdiBusiById()
	this.Data["sdtBdiBusi"] = sdtBdiBusi
	this.TplName = "sdtBdiBusi/sdtBdiBusiDetailConfigIndex.html"
}

//处理类型
func (this *SdtBdiBusiController) ProcessType() {
	type DataType struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	}

	returnData := make([]DataType, 0)
	o := orm.NewOrm()

	var maps []orm.Params = make([]orm.Params, 0)
	_, err := o.Raw(" select process_type as Id, process_name as Text from sdt_bdi_busi_process_type ").Values(&maps)

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		dataType.Id = v["Id"].(string)
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//更新信息
func (this *SdtBdiBusiController) UpdateDetailConfig() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiBusi := new(models.SdtBdiBusi)
	err := this.ParseForm(sdtBdiBusi)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiBusi.UpdateDetailConfig()
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