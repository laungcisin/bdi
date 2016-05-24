package controllers

import (
	"bdi/models"
	"log"
	"strconv"
	"fmt"
)

type SdtBdiResultController struct {
	BaseController
}

// 首页
func (this *SdtBdiResultController) Index() {
	bdiId, _ := this.GetInt("bdiId")
	bdiTypeId, _ := this.GetInt("bdiTypeId")
	this.Data["bdiId"] = bdiId
	this.Data["bdiTypeId"] = bdiTypeId
	this.TplName = "sdtBdiResult/sdtBdiResultIndex.html"
}

/**
根据行数和列数，查询符合条件的数据。
*/
func (this *SdtBdiResultController) All() {
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
	sdtBdiResult := new(models.SdtBdiResult)
	sdtBdiResult.BdiId = bdiId
	sdtBdiResultSlice, num, err := sdtBdiResult.GetAllSdtBdiResult(rows, page)
	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = &sdtBdiResultSlice
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &sdtBdiResultSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 新增Dialog
func (this *SdtBdiResultController) AddPage() {
	bdiId, _ := this.GetInt("bdiId")
	bdiTypeId, _ := this.GetInt("bdiTypeId")
	this.Data["bdiId"] = bdiId
	this.Data["bdiTypeId"] = bdiTypeId
	this.TplName = "sdtBdiResult/addDialog.html"
}

//新增
func (this *SdtBdiResultController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiResult := new(models.SdtBdiResult)
	err := this.ParseForm(sdtBdiResult)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	fmt.Print(sdtBdiResult)
	err = sdtBdiResult.Add()
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
func (this *SdtBdiResultController) UpdatePage() {
	id, err := this.GetInt("id")
	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}
	sdtBdiResult := new(models.SdtBdiResult)
	sdtBdiResult.Id = id
	err = sdtBdiResult.GetSdtBdiResultById()

	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdiResult"] = sdtBdiResult
	this.TplName = "sdtBdiResult/updateDialog.html"
}

//更新信息
func (this *SdtBdiResultController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiResult := new(models.SdtBdiResult)
	err := this.ParseForm(sdtBdiResult)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiResult.Update()
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
func (this *SdtBdiResultController) Delete() {
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

func (this *SdtBdiResultController) DatasourceTree() {
	showLeafLevel, _ := this.GetInt("showLeafLevel")
	this.Data["showLeafLevel"] = showLeafLevel
	this.TplName = "sdtBdiResult/datasourceTreeDialog.html"
}