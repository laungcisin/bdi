package controllers

import (
	"bdi/models"
	"strconv"
	"github.com/astaxie/beego/orm"

)

type SdtPublicDisplayDataController struct {
	BaseController
}

//列出所有的指标库-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiBaseForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	maps, err := models.GetAllSdtBdiBaseForList()
	returnData := []DataType{}

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}


	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的指标域-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiDomainForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	maps, err := models.GetAllSdtBdiDomainForList()
	returnData := []DataType{}

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的指标集-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiSetForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	var sdtBdiSetList = new (models.SdtBdiSet)
	maps, err := sdtBdiSetList.GetAllSdtBdiSetForList()
	returnData := []DataType{}

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的指标保密级别-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiSecrecyForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	returnData := []DataType{}

	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select bdi_secrecy_id as Id, bdi_secrecy_name as Text from sdt_bdi_secrecy ").Values(&maps)

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的指标类型-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiTypeForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	returnData := []DataType{}

	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select bdi_type_id as Id, bdi_type_name as Text from sdt_bdi_type ").Values(&maps)

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}