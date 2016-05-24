package controllers

import (
	"bdi/models"
	"log"
	"strconv"
	"fmt"
	"encoding/json"
	"strings"
)

type SdtBdiBusiRuleController struct {
	BaseController
}

func (this *SdtBdiBusiRuleController) BdiIndex() {
	this.TplName = "sdtBdiBusiRule/sdtBdiIndex.html"
}

func (this *SdtBdiBusiRuleController) Index() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	this.TplName = "sdtBdiBusiRule/sdtBdiBusiRuleIndex.html"
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

	bdiId, _ := this.GetInt("bdiId")
	sdtBdiBusiRule := new(models.SdtBdiBusiRule)
	sdtBdiBusiRule.BdiId = bdiId
	sdtBdiBusiRuleSlice, num, err := sdtBdiBusiRule.GetAllSdtBdiBusiRule(rows, page) //查询指标域所有记录

	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = nil
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &sdtBdiBusiRuleSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

// 选择字段
func (this *SdtBdiBusiRuleController) ColumnSelectTreePage() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	this.TplName = "sdtBdiBusiRule/columnSelectDialog.html"
}

func (this *SdtBdiBusiRuleController)AddColumn()  {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiBusiRule := new(models.SdtBdiBusiRule)
	inserted := this.GetString("inserted")
	updated := this.GetString("updated")
	bdiId, _ := this.GetInt("bdiId")
	fieldsIds := strings.Split(updated, ",")

	var ob []models.TableTreeAttributes
	err = json.Unmarshal([]byte(inserted), &ob)
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if err = sdtBdiBusiRule.AddColumn(bdiId, fieldsIds, ob); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "数据保存失败"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Success = true
	returnData.Message = "数据保存成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return

}

//处理字段
func (this *SdtBdiBusiRuleController) ProcessFields() {
	busiRuleId, _ := this.GetInt("busiRuleId")
	sdtBdiBusiRule := new(models.SdtBdiBusiRule)
	sdtBdiBusiRule.Id = busiRuleId
	sdtBdiBusiRule.GetSdtBdiBusiRuleById()
	this.Data["sdtBdiBusiRule"] = sdtBdiBusiRule
	this.TplName = "sdtBdiBusiRule/processFieldsDialog.html"
}

//处理字段
func (this *SdtBdiBusiRuleController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiBusiRule := new(models.SdtBdiBusiRule)
	err := this.ParseForm(sdtBdiBusiRule)
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiBusiRule.Update()
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