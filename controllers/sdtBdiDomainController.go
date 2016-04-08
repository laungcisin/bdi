package controllers

import (
	"bdi/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type SdtBdiDomainController struct {
	BaseController
}

// 首页
func (this *SdtBdiDomainController) Index() {
	this.TplName = "sdtBdiDomain/sdtBdiDomainIndex.html"
}

/**
根据行数和列数，查询符合条件的指标域的数据。
*/
func (this *SdtBdiDomainController) All() {
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

	sdtBdiDomain := new(models.SdtBdiDomain)
	sdtBdiDomainSlice, num, err := sdtBdiDomain.GetAllSdtBdiDomain(rows, page) //查询指标域所有记录

	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = nil
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &sdtBdiDomainSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 新增Dialog
func (this *SdtBdiDomainController) AddPage() {
	this.TplName = "sdtBdiDomain/addDialog.html"
}

//新增
func (this *SdtBdiDomainController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiDomain := new(models.SdtBdiDomain)
	err := this.ParseForm(sdtBdiDomain)

	if err != nil {
		fmt.Println("参数解析出错！")
		returnData.Success = false
		returnData.Message = "参数解析出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiDomain.Add()//新增

	if err != nil {
		fmt.Println("新增数据出错！")
		returnData.Success = false
		returnData.Message = "新增数据出错！"
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
func (this *SdtBdiDomainController) UpdatePage() {
	bdiDomainId, err := this.GetInt("bdiDomainId")
	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	sdtBdiDomain := new(models.SdtBdiDomain)
	sdtBdiDomain.BdiDomainId = bdiDomainId

	err = sdtBdiDomain.GetSdtBdiDomainById()

	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdiDomain"] = sdtBdiDomain
	this.TplName = "sdtBdiDomain/updateDialog.html"
}

//更新信息
func (this *SdtBdiDomainController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiDomain := new(models.SdtBdiDomain)
	err := this.ParseForm(sdtBdiDomain)

	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiDomain.Update()
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
func (this *SdtBdiDomainController) Delete() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	bdiDomainId, err := this.GetInt("bdiDomainId")
	if err != nil {
		fmt.Println("解析参数出错")
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&models.SdtBdiDomain{BdiDomainId: bdiDomainId}); err != nil {
		fmt.Println("数据删除失败")
		returnData.Success = false
		returnData.Message = "数据删除失败"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Success = true
	returnData.Message = "数据删除成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}
