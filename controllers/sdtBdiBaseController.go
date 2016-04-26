package controllers

import (
	"bdi/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type SdtBdiBaseController struct {
	BaseController
}

// 首页
func (this *SdtBdiBaseController) Index() {
	this.TplName = "sdtBdiBase/sdtBdiBaseIndex.html"
}

/**
列出所有的数据
*/
func (this *SdtBdiBaseController) All() {
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

	sdtBdiBase := new(models.SdtBdiBase)
	sdtBdiBaseSlice, num, err := sdtBdiBase.GetAllSdtBdiBase(rows, page)

	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = nil
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &sdtBdiBaseSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 跳转至新增Dialog页面
func (this *SdtBdiBaseController) AddPage() {
	this.TplName = "sdtBdiBase/addDialog.html"
}

//新增（保存新增的内容）
func (this *SdtBdiBaseController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var sdtBdiBase models.SdtBdiBase
	err := this.ParseForm(&sdtBdiBase)

	if err != nil {
		fmt.Println("解析参数出错！")
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiBase.Add()
	if err != nil {
		fmt.Println("新增出错！")
		returnData.Success = false
		returnData.Message = "新增出错！"
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

// 跳转至编辑Dialog页面
func (this *SdtBdiBaseController) UpdatePage() {
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println("解析参数出错！")
		return
	}

	sdtBdiBase := new(models.SdtBdiBase)
	sdtBdiBase.Id = id

	sdtBdiBase.GetSdtBdiBaseById()

	if err != nil {
		fmt.Println("查詢出错！")
		return
	}

	this.Data["sdtBdiBase"] = sdtBdiBase
	this.TplName = "sdtBdiBase/updateDialog.html"
}

//更新信息
func (this *SdtBdiBaseController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var sdtBdiBase models.SdtBdiBase
	err := this.ParseForm(&sdtBdiBase)

	if err != nil {
		fmt.Println("解析参数出错！")
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiBase.Update()
	if err != nil {
		fmt.Println("更新出错！")
		returnData.Success = false
		returnData.Message = "更新出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Success = true
	returnData.Message = "更新数据成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//删除
func (this *SdtBdiBaseController) Delete() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	bdiBaseId, err := this.GetInt("bdiBaseId")
	fmt.Println("bdiBaseId: ", bdiBaseId)
	if err != nil {
		fmt.Println("解析参数出错")
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	if _, err := o.Delete(&models.SdtBdiBase{Id: bdiBaseId}); err != nil {
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
