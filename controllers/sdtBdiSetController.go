package controllers

import (
	"bdi/models"
	"log"
	"strconv"
)

type SdtBdiSetController struct {
	BaseController
}

// 首页
func (this *SdtBdiSetController) Index() {
	this.TplName = "sdtBdiSet/sdtBdiSetIndex.html"
}

/**
根据行数和列数，查询符合条件的指标集的数据。
*/
func (this *SdtBdiSetController) All() {
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

	sdtBdiSet := new(models.SdtBdiSet)
	sdtBdiSetSlice, num, err := sdtBdiSet.GetAllSdtBdiSet(rows, page)

	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = nil
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &sdtBdiSetSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 新增Dialog
func (this *SdtBdiSetController) AddPage() {
	this.TplName = "sdtBdiSet/addDialog.html"
}

//新增
func (this *SdtBdiSetController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var sdtBdiSet models.SdtBdiSet
	err := this.ParseForm(&sdtBdiSet)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiSet.Add()
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
func (this *SdtBdiSetController) UpdatePage() {
	var err error
	id, err := this.GetInt("id")

	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	sdtBdiSet := new(models.SdtBdiSet)
	sdtBdiSet.Id = id
	err = sdtBdiSet.GetSdtBdiSetById()

	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdiSet"] = sdtBdiSet
	this.TplName = "sdtBdiSet/updateDialog.html"
}

//更新信息
func (this *SdtBdiSetController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var sdtBdiSet models.SdtBdiSet
	err := this.ParseForm(&sdtBdiSet)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiSet.Update()
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
func (this *SdtBdiSetController) Delete() {
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
