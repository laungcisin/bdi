package controllers

import (
	"bdi/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	//	"time"
	"encoding/json"
	"time"
)

type SdtBdiBaseController struct {
	BaseController
}

// 首页
func (this *SdtBdiBaseController) Index() {
	//	o := orm.NewOrm()
	//	var maps []orm.Params
	//	var s = "SELECT Id, Bdi_Base_Code, Bdi_Base_Name, Bdi_Base_Year, Bus_date, Tot_Name, Abb_Name, Syn_Name, En_Tot_Name, En_Name, En_Abb_Name, Abb_Code, Std_ID, Explains, Remarks, Bdi_Domain_Flag_Code, Bdi_Domain_Code_Method, Bdi_Set_Code_Method, Bdi_Type_Code_Method, Bdi_Code_Method, Adm_Div_Code, Unit_Code, Manager_User_Code, Pause_Flag_Code, Used_Flag_Code, Code_Lvl_No, Bot_Flag_Code, User_Code, User_Name, Create_Time, Start_Use_Date FROM sdt_bdi_base where Id = ?"
	//
	//	o.Raw(s, 2).Values(&maps)
	//	fmt.Println("结果1：", maps)
	//
	//	var sdtBdiBase models.SdtBdiBase
	//	o.Raw(s, 2).QueryRow(&sdtBdiBase)
	//	fmt.Println("结果2：", sdtBdiBase)

	// o := orm.NewOrm()
	// var sdtBdiBase models.SdtBdiBase
	// err := o.QueryTable(sdtBdiBase.TableName()).Filter("id", 2).One(&sdtBdiBase)
	// if err == orm.ErrMultiRows {
	// 	// 多条的时候报错
	// 	fmt.Printf("Returned Multi Rows Not One")
	// }
	// if err == orm.ErrNoRows {
	// 	// 没有找到记录
	// 	fmt.Printf("Not row found")
	// }

	// fmt.Println("结果1：", sdtBdiBase)

	// var sdtBdiBaseList []*models.SdtBdiBase
	// num, err := o.QueryTable(sdtBdiBase.TableName()).Limit(-1).All(&sdtBdiBaseList)
	// if err != nil {
	// 	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	// }
	// fmt.Println("结果2：", len(sdtBdiBaseList))

	// this.Data["pageTitle"] = "系统概况"
	this.TplName = "sdtBdiBase/sdtBdiBaseIndex.html"
}

/**
列出所有的数据
*/
func (this *SdtBdiBaseController) All() {
	rows, _ := strconv.Atoi(this.GetString("rows"))
	page, _ := strconv.Atoi(this.GetString("page"))

	var sdtBdiBaseList []*models.SdtBdiBase
	var sdtBdiBase models.SdtBdiBase

	if _, err := orm.NewOrm().QueryTable(sdtBdiBase.TableName()).Limit(page*rows, (page-1)*rows).All(&sdtBdiBaseList); err != nil {
		return
	}

	num, err := orm.NewOrm().QueryTable(sdtBdiBase.TableName()).Limit(-1).Count()
	if err != nil {
		return
	}

	//返回的json数据
	returnData := struct {
		Total interface{} `json:"total"`
		Rows  interface{} `json:"rows"`
	}{}

	returnData.Total = num
	returnData.Rows = &sdtBdiBaseList
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
}

//编辑首页
func (this *SdtBdiBaseController) Content() {
	bdiBaseId, err := this.GetInt("bdiBaseId")
	var sdtBdiBase models.SdtBdiBase
	if err != nil {
		fmt.Println("参数解析出错！")
		return
	}

	orm.NewOrm().QueryTable(sdtBdiBase.TableName()).Filter("BdiBaseId", bdiBaseId).One(&sdtBdiBase)
	this.Data["sdtBdiBase"] = sdtBdiBase
	this.TplName = "sdtBdiBase/sdtBdiBaseContent.html"
}

//基础信息
func (this *SdtBdiBaseController) BaseInfo() {
	id, err := this.GetInt("id")
	var sdtBdiBase models.SdtBdiBase

	if err == nil {
		orm.NewOrm().QueryTable(sdtBdiBase.TableName()).Filter("id", id).One(&sdtBdiBase)
	}

	this.Data["sdtBdiBase"] = sdtBdiBase
	this.TplName = "sdtBdiBase/properties/baseInfo.html"
}

//更新信息
func (this *SdtBdiBaseController) Update() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var sdtBdiBaseSlice []models.SdtBdiBase
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &sdtBdiBaseSlice)

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
	for _, v := range sdtBdiBaseSlice {
		if _, err = o.Update(&v); err != nil {
			o.Rollback()
			fmt.Println("数据库保存出错！")
			returnData.Success = false
			returnData.Message = "数据库保存出错！"
			this.Data[JSON_STRING] = returnData
			this.ServeJSON()
			return
		}
	}
	// 此过程中的所有使用 o Ormer 对象的查询都在事务处理范围内
	err = o.Commit()

	//
	//	if err := this.ParseForm(&sdtBdiBase); err != nil {
	//		fmt.Println("解析参数出错")
	//		returnData.Success = false
	//		returnData.Message = "解析参数出错"
	//		this.Data[JSON_STRING] = returnData
	//		this.ServeJSON()
	//		return
	//	}
	//
	//	//日期处理
	//	createTime := this.GetString("CreateTime")
	//	if theTime, err := time.Parse(GOLANG_TIME_FORMAT, createTime); err == nil {
	//		sdtBdiBase.CreateTime = theTime
	//	}
	//
	//	if _, err := orm.NewOrm().Update(&sdtBdiBase); err != nil {
	//		fmt.Println("保存出错")
	//		returnData.Success = false
	//		returnData.Message = "保存出错"
	//		this.Data[JSON_STRING] = returnData
	//		this.ServeJSON()
	//		return
	//	}
	//
	//	//必须指定是json，返回的数据在returnData中定义。
	returnData.Success = true
	returnData.Message = "数据更新成功！"
	this.Data[JSON_STRING] = returnData
	this.ServeJSON()
	return
}

//新增
func (this *SdtBdiBaseController) Add() {
	returnData := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	var sdtBdiBase models.SdtBdiBase
	err := this.ParseForm(&sdtBdiBase)
	if err != nil {
		fmt.Println("解析参数出错")
		returnData.Success = false
		returnData.Message = "解析参数出错"
		this.Data[JSON_STRING] = returnData
		this.ServeJSON()
		return
	}

	sdtBdiBase.CreateTime = time.Now()
	if _, err := orm.NewOrm().Insert(&sdtBdiBase); err != nil {
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
	if _, err := o.Delete(&models.SdtBdiBase{BdiBaseId: bdiBaseId}); err != nil {
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
