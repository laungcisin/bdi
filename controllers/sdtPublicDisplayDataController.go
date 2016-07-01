package controllers

import (
	"bdi/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type SdtPublicDisplayDataController struct {
	BaseController
}

//列出所有的指标库-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiBaseForList() {
	 type DataType struct {
	 	Id   int    `json:"id"`
	 	Text string `json:"text"`
	 }

	 maps, err := models.GetAllSdtBdiBaseForList()
	 returnData := []DataType{}

	 if err != nil {
	 	this.Data[JSON_STRING] = returnData
	 	this.ServeJSON()
	 	return
	 }

	 for _, v := range maps {
	 	dataType := new(DataType)
	 	id, _ := strconv.Atoi(v["Id"].(string))
	 	dataType.Id = id
	 	dataType.Text = v["Text"].(string)
	 	returnData = append(returnData, *dataType)
	 }

	 this.Data[JSON_STRING] = returnData
	 this.ServeJSON()
	return
}

//列出所有的指标域-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiDomainForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	maps, err := models.GetAllSdtBdiDomainForList()
	returnData := []DataType{}

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的指标集-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiSetForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	var sdtBdiSetList = new(models.SdtBdiSet)
	maps, err := sdtBdiSetList.GetAllSdtBdiSetForList()
	returnData := []DataType{}

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的指标保密级别-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiSecrecyForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	returnData := []DataType{}

	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select bdi_secrecy_id as Id, bdi_secrecy_name as Text from sdt_bdi_secrecy ").Values(&maps)

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的指标类型-用于下拉列表
func (this *SdtPublicDisplayDataController) SdtBdiTypeForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	returnData := []DataType{}

	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select id as Id, type_name as Text from sdt_bdi_type ").Values(&maps)

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的数据源类型-用于下拉列表
func (this *SdtPublicDisplayDataController) DatasourceTypeForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	returnData := make([]DataType, 0)
	o := orm.NewOrm()

	var maps []orm.Params = make([]orm.Params, 0)
	_, err := o.Raw(" select t.id as Id, t.type_name as Text from sdt_bdi_datasource_type t ").Values(&maps)

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的数据源-用于下拉列表
func (this *SdtPublicDisplayDataController) DatasourceForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	returnData := make([]DataType, 0)
	o := orm.NewOrm()

	var maps []orm.Params = make([]orm.Params, 0)
	_, err := o.Raw(" select t.id as Id, t.name as Text from sdt_bdi_datasource t ").Values(&maps)

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//列出所有的Stat-Dim-用于下拉列表
func (this *SdtPublicDisplayDataController) StatDimForList() {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	returnData := []DataType{}

	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select bdi_stat_dim_id as Id, bdi_stat_dim_name as Text from sdt_bdi_stat_dim ").Values(&maps)

	if err != nil {
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

func (this *SdtPublicDisplayDataController) ProcessTypeForListByCondition() {
	mainClass, _ := this.GetInt("mainClass")
	subClass, _ := this.GetInt("subClass")

	type DataType struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	}

	returnData := []DataType{}

	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select t.process_type as Id, t.process_name as Text from sdt_bdi_process_type t where t.main_class = ? and t.sub_class = ? ", mainClass, subClass).Values(&maps)

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

func (this *SdtPublicDisplayDataController) ProcessTypeForListAll() {
	type DataType struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	}

	returnData := []DataType{}

	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select t.process_type as Id, concat(t.process_name, '(', t.process_type, ')') as Text from sdt_bdi_process_type t ").Values(&maps)

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
