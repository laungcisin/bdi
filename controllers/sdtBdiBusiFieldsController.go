package controllers

import (
	"bdi/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type SdtBdiBusiFieldsController struct {
	BaseController
}

// 首页
func (this *SdtBdiBusiFieldsController) Index() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	this.TplName = "sdtBdiBusiFields/sdtBdiBusiFieldsIndex.html"
}

/**
根据行数和列数，查询符合条件的数据。
*/
func (this *SdtBdiBusiFieldsController) All() {
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
	sdtBdiBusiFields := new(models.SdtBdiBusiFields)
	sdtBdiBusiFields.BdiId = bdiId
	sdtBdiBusiSlice, num, err := sdtBdiBusiFields.GetAllSdtBdiBusiFields(rows, page)
	if err != nil {
		log.Fatal("查询数据失败！")
		returnData.Total = 0
		returnData.Rows = &sdtBdiBusiSlice
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	returnData.Total = num
	returnData.Rows = &sdtBdiBusiSlice
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//新增字段首页
func (this *SdtBdiBusiFieldsController) AddFields() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	this.TplName = "sdtBdiBusiFields/addFieldsDialog.html"
}

//处理字段
func (this *SdtBdiBusiFieldsController) ProcessType() {
	type DataType struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	}

	returnData := make([]DataType, 0)
	o := orm.NewOrm()

	var maps []orm.Params = make([]orm.Params, 0)
	_, err := o.Raw(" select process_type as Id, process_name as Text from sdt_bdi_process_type ").Values(&maps)

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		dataType.Id = v["Id"].(string)
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//处理字段
func (this *SdtBdiBusiFieldsController) ProcessFields() {
	busiFieldsId, _ := this.GetInt("busiFieldsId")
	sdtBdiBusiFields := new(models.SdtBdiBusiFields)
	sdtBdiBusiFields.Id = busiFieldsId
	sdtBdiBusiFields.GetSdtBdiBusiById()
	this.Data["sdtBdiBusiFields"] = sdtBdiBusiFields
	this.TplName = "sdtBdiBusiFields/processFieldsDialog.html"
}

//字段加工
func (this *SdtBdiBusiFieldsController) ProcessFieldsAdd() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiBusiFields := new(models.SdtBdiBusiFields)
	err = this.ParseForm(sdtBdiBusiFields)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	busiFieldsId, _ := this.GetInt("busiFieldsId")
	sdtBdiBusiFields.Id = busiFieldsId
	err = sdtBdiBusiFields.UpdateFields()
	if err != nil {
		returnData.Success = false
		returnData.Message = "更新数据出错"
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

// 新增Dialog
func (this *SdtBdiBusiFieldsController) AddPage() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	this.TplName = "sdtBdiBusiFields/addDialog.html"
}

//新增
func (this *SdtBdiBusiFieldsController) Add() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var ob []models.TableTreeAttributes
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	sdtBdiBusi := new(models.SdtBdiBusi)
	err = sdtBdiBusi.AddBusiAndAddField(ob)
	if err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "保存数据出错"
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
func (this *SdtBdiBusiFieldsController) UpdatePage() {
	id, err := this.GetInt("id")
	if err != nil {
		log.Fatal("解析参数出错！")
		return
	}
	sdtBdiBusi := new(models.SdtBdiBusi)
	sdtBdiBusi.Id = id
	err = sdtBdiBusi.GetSdtBdiBusiById()

	if err != nil {
		fmt.Println(err)
		log.Fatal("解析参数出错！")
		return
	}

	this.Data["sdtBdiBusi"] = sdtBdiBusi
	this.TplName = "sdtBdiBusi/updateDialog.html"
}

//更新信息
func (this *SdtBdiBusiFieldsController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiBusi := new(models.SdtBdiBusi)
	err := this.ParseForm(sdtBdiBusi)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiBusi.Update()
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

//行上移
func (this *SdtBdiBusiFieldsController) RowMoveUp() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiBusiFields := new(models.SdtBdiBusiFields)
	updated := this.GetString("updated")

	if err = json.Unmarshal([]byte(updated), sdtBdiBusiFields); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if err = sdtBdiBusiFields.RowMoveUp(); err != nil {
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

//行下移
func (this *SdtBdiBusiFieldsController) RowMoveDown() {
	var err error
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiBusiFields := new(models.SdtBdiBusiFields)
	updated := this.GetString("updated")

	if err = json.Unmarshal([]byte(updated), sdtBdiBusiFields); err != nil {
		fmt.Println(err)
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	if err = sdtBdiBusiFields.RowMoveDown(); err != nil {
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