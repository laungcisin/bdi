package controllers

import (
	"bdi/models"
	"log"
	"strconv"
)

type SdtBdiCfgKpiController struct {
	BaseController
}

// Kpi首页
func (this *SdtBdiCfgKpiController) Index() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId

	this.TplName = "sdtBdiCfgKpi/sdtBdiCfgKpiIndex.html"
}

/**
根据行数和列数，查询符合条件的指标集的数据。
*/
func (this *SdtBdiCfgKpiController) All() {
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

	sdtBdiCfgKpi := new(models.SdtBdiCfgKpi)
	sdtBdiCfgKpiSetSlice, num, err := sdtBdiCfgKpi.GetAllSdtBdiCfgKpi(rows, page) //查询所有

	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = nil
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &sdtBdiCfgKpiSetSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 新增Dialog
func (this *SdtBdiCfgKpiController) AddPage() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	this.TplName = "sdtBdiCfgKpi/addDialog.html"
}

//新增
func (this *SdtBdiCfgKpiController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiCfgKpi := new(models.SdtBdiCfgKpi)
	err := this.ParseForm(sdtBdiCfgKpi)

	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiCfgKpi.Add()
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
func (this *SdtBdiCfgKpiController) UpdatePage() {
	kpiId, err := this.GetInt("kpiId")
	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}
	sdtBdiCfgKpi := new(models.SdtBdiCfgKpi)
	sdtBdiCfgKpi.KpiId = kpiId
	err = sdtBdiCfgKpi.GetSdtBdiCfgKpiById()

	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdiCfgKpi"] = sdtBdiCfgKpi
	this.TplName = "sdtBdiCfgKpi/updateDialog.html"
}

//更新信息
func (this *SdtBdiCfgKpiController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiCfgKpi := new(models.SdtBdiCfgKpi)
	err := this.ParseForm(sdtBdiCfgKpi)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiCfgKpi.Update()
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
func (this *SdtBdiCfgKpiController) Delete() {
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
