package controllers

import (
	"bdi/models"
	"fmt"
	"strconv"
)

type SdtBdiBusiConfigController struct {
	BaseController
}

// 首页
func (this *SdtBdiBusiConfigController) Index() {
	bdiId, _ := this.GetInt("bdiId")
	this.Data["bdiId"] = bdiId
	this.TplName = "sdtBdiBusiConfig/sdtBdiBusiConfigIndex.html"
}

/**
根据行数和列数，查询符合条件的数据。
*/
func (this *SdtBdiBusiConfigController) All() {
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
	sdtBdiBusiConfig := new(models.SdtBdiBusiConfig)
	sdtBdiBusiConfig.BdiId = bdiId
	sdtBdiBusiSlice, num, err := sdtBdiBusiConfig.GetAllSdtBdiBusiConfig(rows, page)
	if err != nil {
		fmt.Println("查询数据失败！")
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

// 更新Dialog
func (this *SdtBdiBusiConfigController) UpdatePage() {
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println("解析参数出错！")

		return
	}
	sdtBdiBusiConfig := new(models.SdtBdiBusiConfig)
	sdtBdiBusiConfig.Id = id
	err = sdtBdiBusiConfig.GetSdtBdiBusiConfigById()

	if err != nil {
		fmt.Println(err)
		fmt.Println("解析参数出错！")
		return
	}

	this.Data["sdtBdiBusiConfig"] = sdtBdiBusiConfig
	this.TplName = "sdtBdiBusiConfig/updateDialog.html"
}

//更新信息
func (this *SdtBdiBusiConfigController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiBusiConfig := new(models.SdtBdiBusiConfig)
	err := this.ParseForm(sdtBdiBusiConfig)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	fmt.Println("sdtBdiBusiConfig: ", sdtBdiBusiConfig)

	err = sdtBdiBusiConfig.Update()
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

// 首页
func (this *SdtBdiBusiConfigController) ColumnSelectTreePage() {
	this.TplName = "sdtBdiBusiConfig/columnTreeDialog.html"
}

//同步表信息
func (this *SdtBdiBusiConfigController) Synchronize() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	sdtBdiBusiConfig := new(models.SdtBdiBusiConfig)
	err := this.ParseForm(sdtBdiBusiConfig)
	if err != nil {
		returnData.Success = false
		returnData.Message = "解析参数出错！"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	err = sdtBdiBusiConfig.Synchronize()
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
