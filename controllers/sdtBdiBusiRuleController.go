package controllers

import (
	"bdi/models"
	"log"
	"strconv"
)

type SdtBdiBusiRuleController struct {
	BaseController
}

// 首页
func (this *SdtBdiBusiRuleController) BdiIndex() {
	this.TplName = "sdtBdiBusiRule/sdtBdiIndex.html"
}

/**
根据行数和列数，查询符合条件的数据。
*/
func (this *SdtBdiBusiRuleController) All() {
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

	sdtBdi := new(models.SdtBdi)
	sdtBdiSetSlice, num, err := sdtBdi.GetAllSdtBdi(rows, page) //查询指标域所有记录

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

// 选择字段
func (this *SdtBdiBusiRuleController) ColumnSelectTreePage() {
	this.TplName = "sdtBdiBusiRule/columnSelectDialog.html"
}