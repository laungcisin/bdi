package controllers

import (
	"bdi/models"
	"fmt"
	"log"
	"strconv"
)

type SdtBdiRuleController struct {
	BaseController
}

// 首页
func (this *SdtBdiRuleController) Index() {
	bdiRuleSetId, _ := this.GetInt("bdiRuleSetId")
	Id, _ := this.GetInt("Id")
	this.Data["bdiRuleSetId"] = bdiRuleSetId
	this.Data["Id"] = Id
	this.TplName = "sdtBdiRule/sdtBdiRuleIndex.html"
}

/**
根据行数和列数，查询符合条件的规则数据。
*/
func (this *SdtBdiRuleController) All() {
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
		rows = 100
	}

	//页数
	page, err = strconv.Atoi(this.GetString("page"))
	if err != nil {
		page = 1
	}

	bdiRuleSetId, _ := this.GetInt("bdiRuleSetId")
	sdtBdiRule := new(models.SdtBdiRule)
	sdtBdiRule.BdiRuleSetId = bdiRuleSetId
	sdtBdiRuleSlice, num, err := sdtBdiRule.GetAllSdtBdiRuleByRuleSetId(rows, page) //查询指标域所有记录

	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = 0
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = sdtBdiRuleSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 新增Dialog
func (this *SdtBdiRuleController) AddPage() {
	Id, _ := this.GetInt("Id")
	bdiRuleSetId, _ := this.GetInt("bdiRuleSetId")
	this.Data["Id"] = Id
	this.Data["bdiRuleSetId"] = bdiRuleSetId
	this.TplName = "sdtBdiRule/addDialog.html"
}

//新增
func (this *SdtBdiRuleController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiRule := new(models.SdtBdiRule)
	err := this.ParseForm(sdtBdiRule)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiRule.Add()
	if err != nil {
		fmt.Println("err: ", err)
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
func (this *SdtBdiRuleController) UpdatePage() {
	Id, err := this.GetInt("Id")
	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}
	sdtBdiRule := new(models.SdtBdi)
	sdtBdiRule.Id = Id
	err = sdtBdiRule.GetSdtBdiById()

	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdiRule"] = sdtBdiRule
	this.TplName = "sdtBdiRule/updateDialog.html"
}

//更新信息
func (this *SdtBdiRuleController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiRule := new(models.SdtBdi)
	err := this.ParseForm(sdtBdiRule)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiRule.Update()
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
