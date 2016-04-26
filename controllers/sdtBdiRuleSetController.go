package controllers

import (
	"bdi/models"
	"encoding/json"
	//	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	//	"fmt"
	"fmt"
)

type SdtBdiRuleSetController struct {
	BaseController
}

// 首页
func (this *SdtBdiRuleSetController) Index() {
	bdiId, _ := this.GetInt("bdiId")
	sdtBdi := new(models.SdtBdi)
	sdtBdi.Id = bdiId
	sdtBdi.GetSdtBdiById()

	this.Data["sdtBdi"] = sdtBdi
	this.TplName = "sdtBdiRuleSet/sdtBdiRuleSetIndex.html"
}

/**
根据行数和列数，查询符合条件的规则集数据。
*/
func (this *SdtBdiRuleSetController) All() {
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
		rows = 100
	}

	//页数
	page, err = strconv.Atoi(this.GetString("page"))
	if err != nil {
		page = 1
	}

	bdiId, _ := this.GetInt("bdiId")

	sdtBdiRuleSet := new(models.SdtBdiRuleSet)
	sdtBdiRuleSet.BdiId = bdiId
	sdtBdiRuleSetSlice, num, err := sdtBdiRuleSet.GetAllSdtBdiRuleSetByBdiId(rows, page) //查询指标域所有记录

	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = 0
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = sdtBdiRuleSetSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//新增
func (this *SdtBdiRuleSetController) Add() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiRuleSet := new(models.SdtBdiRuleSet)
	inserted := this.GetString("inserted")

	if err = json.Unmarshal([]byte(inserted), sdtBdiRuleSet); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if err = sdtBdiRuleSet.Add(); err != nil {
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

//更新
func (this *SdtBdiRuleSetController) Update() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiRuleSet := new(models.SdtBdiRuleSet)
	updated := this.GetString("updated")

	if err = json.Unmarshal([]byte(updated), sdtBdiRuleSet); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if err = sdtBdiRuleSet.Update(); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "数据更新失败"
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

//行上移
func (this *SdtBdiRuleSetController) RowMoveUp() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiRuleSet := new(models.SdtBdiRuleSet)
	updated := this.GetString("updated")

	if err = json.Unmarshal([]byte(updated), sdtBdiRuleSet); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if err = sdtBdiRuleSet.RowMoveUp(); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "数据更新失败"
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

//行上移
func (this *SdtBdiRuleSetController) RowMoveDown() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiRuleSet := new(models.SdtBdiRuleSet)
	updated := this.GetString("updated")

	if err = json.Unmarshal([]byte(updated), sdtBdiRuleSet); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if err = sdtBdiRuleSet.RowMoveDown(); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "数据更新失败"
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
func (this *SdtBdiRuleSetController) Delete() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiRuleSet := new(models.SdtBdiRuleSet)
	err = this.ParseForm(sdtBdiRuleSet)
	if err != nil {
		fmt.Println("解析参数出错")
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if err = sdtBdiRuleSet.Delete(); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "数据更新失败"
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

//批量更新/新增/删除
func (this *SdtBdiRuleSetController) BatchAdd() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	inserted := this.GetString("inserted")
	deleted := this.GetString("deleted")
	updated := this.GetString("updated")

	var insertedRuleSetSlice []models.SdtBdiRuleSet
	var deletedRuleSetSlice []models.SdtBdiRuleSet
	var updatedRuleSetSlice []models.SdtBdiRuleSet

	if len(inserted) > 0 || inserted != "" {
		if err = json.Unmarshal([]byte(inserted), &insertedRuleSetSlice); err != nil {
			fmt.Println(err)
			returnData.Success = false
			returnData.Message = "解析参数出错"
			this.Data[JSON_STRING] = returnData
			this.ServeJSON()
			return
		}
	}

	if len(deleted) > 0 || deleted != "" {
		if err = json.Unmarshal([]byte(deleted), &deletedRuleSetSlice); err != nil {
			fmt.Println(err)
			returnData.Success = false
			returnData.Message = "解析参数出错"
			this.Data[JSON_STRING] = returnData
			this.ServeJSON()
			return
		}
	}

	if len(updated) > 0 || updated != "" {
		if err = json.Unmarshal([]byte(updated), &updatedRuleSetSlice); err != nil {
			fmt.Println(err)
			returnData.Success = false
			returnData.Message = "解析参数出错"
			this.Data[JSON_STRING] = returnData
			this.ServeJSON()
			return
		}
	}

	sdtBdiRuleSet := new(models.SdtBdiRuleSet)
	if err = sdtBdiRuleSet.UpdateStatus(insertedRuleSetSlice, deletedRuleSetSlice, updatedRuleSetSlice); err != nil {
		returnData.Success = false
		returnData.Message = "数据更新出错"
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

// 更新Dialog
func (this *SdtBdiRuleSetController) UpdatePage() {
	bdiSetId, err := this.GetInt("bdiSetId")

	sdtBdiSet := new(models.SdtBdiSet)
	sdtBdiSet.Id = bdiSetId
	err = sdtBdiSet.GetSdtBdiSetById()

	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdiSet"] = sdtBdiSet
	this.TplName = "sdtBdiSet/updateDialog.html"
}
