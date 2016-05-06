package controllers

import (
	"bdi/models"
	"log"
	"strconv"
)

type SdtBdiResultFieldsController struct {
	BaseController
}

// 首页
func (this *SdtBdiResultFieldsController) Index() {
	resultId, _ := this.GetInt("resultId")
	this.Data["resultId"] = resultId
	this.TplName = "sdtBdiResultFields/sdtBdiResultFieldsIndex.html"
}

/**
根据行数和列数，查询符合条件的数据。
*/
func (this *SdtBdiResultFieldsController) All() {
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

	sdtBdiResultFields := new(models.SdtBdiResultFields)
	sdtBdiResultFieldsSlice, num, err := sdtBdiResultFields.GetAllSdtBdiResultFields(rows, page)
	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = &sdtBdiResultFieldsSlice
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &sdtBdiResultFieldsSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 新增Dialog
func (this *SdtBdiResultFieldsController) AddPage() {
	this.TplName = "sdtBdiResultFields/addDialog.html"
}

//新增
func (this *SdtBdiResultFieldsController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiResultFields := new(models.SdtBdiResultFields)
	err := this.ParseForm(sdtBdiResultFields)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiResultFields.Add()
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
func (this *SdtBdiResultFieldsController) UpdatePage() {
	id, err := this.GetInt("id")
	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}
	sdtBdiResultFields := new(models.SdtBdiResultFields)
	sdtBdiResultFields.Id = id
	err = sdtBdiResultFields.GetSdtBdiResultFieldsById()

	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdiResultFields"] = sdtBdiResultFields
	this.TplName = "sdtBdiResultFields/updateDialog.html"
}

//更新信息
func (this *SdtBdiResultFieldsController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiResultFields := new(models.SdtBdiResultFields)
	err := this.ParseForm(sdtBdiResultFields)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiResultFields.Update()
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
func (this *SdtBdiResultFieldsController) Delete() {
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
