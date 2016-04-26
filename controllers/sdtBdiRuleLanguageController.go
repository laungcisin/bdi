package controllers

import (
	"bdi/models"
	"log"
)

type SdtBdiRuleLanguageController struct {
	BaseController
}

// 首页
func (this *SdtBdiRuleLanguageController) Index() {
	Id, _ := this.GetInt("Id")
	this.Data["Id"] = Id
	this.TplName = "sdtBdiRuleLanguage/sdtBdiRuleIndex.html"
}

/**
根据行数和列数，查询符合条件的规则数据。
*/
func (this *SdtBdiRuleLanguageController) All() {
	language := new(models.SdtBdiRuleLanguage)
	returnData, _ := language.All()
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 新增Dialog
func (this *SdtBdiRuleLanguageController) AddPage() {
	Id, _ := this.GetInt("Id")
	this.Data["Id"] = Id
	this.TplName = "sdtBdiRuleLanguage/addDialog.html"
}

//新增
func (this *SdtBdiRuleLanguageController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	language := new(models.SdtBdiRuleLanguage)
	err := this.ParseForm(language)

	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if isExist, err := language.NameIsExist(); isExist || err != nil {
		returnData.Success = false
		returnData.Message = "名称已存在，请修改名称"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = language.Add()
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
func (this *SdtBdiRuleLanguageController) UpdatePage() {
	Id, err := this.GetInt("Id")
	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}
	sdtBdi := new(models.SdtBdi)
	sdtBdi.Id = Id
	err = sdtBdi.GetSdtBdiById()

	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdi"] = sdtBdi
	this.TplName = "sdtBdiRuleLanguage/updateDialog.html"
}

//更新信息
func (this *SdtBdiRuleLanguageController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdi := new(models.SdtBdi)
	err := this.ParseForm(sdtBdi)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdi.Update()
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
