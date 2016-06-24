package controllers

import (
	"bdi/models"
)

type SdtBdiTypeController struct {
	BaseController
}

// 新增Dialog
func (this *SdtBdiTypeController) AddPage() {
	this.TplName = "sdtBdiType/addDialog.html"
}

//新增
func (this *SdtBdiTypeController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiType := new(models.SdtBdiType)
	err := this.ParseForm(sdtBdiType)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiType.Add()
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
