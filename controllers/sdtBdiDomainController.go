package controllers

import (
	"bdi/models"
	"encoding/json"
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

	sdtBdiDomainSlice, num, err := models.GetAllSdtBdiDomain(rows, page) //查询指标域所有记录

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

//更新信息
func (this *SdtBdiDomainController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var sdtBdiDomainSlice []models.SdtBdiDomain
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &sdtBdiDomainSlice)

	if err != nil {
		fmt.Println("参数解析出错: ", err)
		returnData.Success = false
		returnData.Message = "参数解析出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		fmt.Println("数据库连接出错！")
		returnData.Success = false
		returnData.Message = "数据库连接出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	// 事务处理过程
	for _, v := range sdtBdiDomainSlice {
		v.BdiBaseName = "" //置空，ORM框架使用的问题
		if _, err = o.Update(&v); err != nil {
			o.Rollback()
			fmt.Println("数据库更新出错！")
			returnData.Success = false
			returnData.Message = "数据库更新出错！"
			this.Data[JSON_STRING] = returnData
			this.ServeJSON()
			return
		}
	}
	// 此过程中的所有使用 o Ormer 对象的查询都在事务处理范围内
	err = o.Commit()

	returnData.Success = true
	returnData.Message = "数据更新成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//新增
func (this *SdtBdiDomainController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	bdiBaseId, err := this.GetInt("bdiBaseId")
	if err != nil {
		fmt.Println("参数解析出错！")
		returnData.Success = false
		returnData.Message = "参数解析出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	bdiDomainName := this.GetString("bdiDomainName")
	adtFlag := this.GetString("adtFlag")
	remarks := this.GetString("remarks")

	sdtBdiDomain := new(models.SdtBdiDomain)
	sdtBdiDomain.BdiBaseId = bdiBaseId
	sdtBdiDomain.BdiBaseName = bdiDomainName
	sdtBdiDomain.AdtFlag = adtFlag
	sdtBdiDomain.Remarks = remarks
	err = models.AddSdtBdiDomain(sdtBdiDomain)//新增

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

//删除
func (this *SdtBdiDomainController) Delete() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	bdiDomainId, err := this.GetInt("bdiDomainId")
	fmt.Println("bdiDomainId: ", bdiDomainId)
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
