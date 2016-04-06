package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)

//规则语言
type SdtBdiRuleLanguage struct {
	BdiRuleLanguageId int    `form:"bdiRuleLanguageId"` //主键
	BdiId             int    `form:"bdiId"`             //指标项Id
	BdiBusinessDataId int    `form:"bdiBusinessDataId"` //业务数据Id
	LanguageName      string `form:"languageName"`      //名称
	LanguageNum       int    `form:"languageNum"`       //序号
	DefaultValue      string `form:"defaultValue"`      //默认值
	Remarks           string `form:"remarks"`           //备注

	CreateUserId int       //创建人ID
	CreateTime   time.Time //创建时间
	ModifyUserId int       //修改人ID
	ModifyTime   time.Time //修改时间
}

func (u *SdtBdiRuleLanguage) TableName() string {
	return "sdt_bdi_rule_language"
}

func (this *SdtBdiRuleLanguage) All()(interface{}, error) {
	type DataType struct {
		Id   int    `json:"id"`
		Text string `json:"text"`
	}

	returnData := []DataType{}

	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select bdi_rule_language_id as Id, language_name as Text from sdt_bdi_rule_language ").Values(&maps)

	if err != nil {
		return nil, err
	}

	for _, v := range maps {
		dataType := new(DataType)
		id, _ := strconv.Atoi(v["Id"].(string))
		dataType.Id = id
		dataType.Text = v["Text"].(string)
		returnData = append(returnData, *dataType)
	}

	return returnData, nil
}

//检测规则语言是否存在
func (this *SdtBdiRuleLanguage)NameIsExist()(bool, error)  {
	var err error
	o := orm.NewOrm()
	var querySql string = " select count(*) from sdt_bdi_rule_language where bdi_id = ? and language_name = ? "
	var languageNum int64 = 0
	err = o.Raw(querySql, this.BdiId, this.LanguageName).QueryRow(&languageNum)

	if err != nil {
		return true, err
	}

	if languageNum > 0 {
		return true, nil
	}
	return false, nil
}

//新增
func (this *SdtBdiRuleLanguage) Add() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	//新增前获取最大RuleSetLevel
	var querySql string = " select max(language_num) as language_num from sdt_bdi_rule_language "
	languageNum := new(int)
	err = o.Raw(querySql).QueryRow(languageNum)

	if err != nil {
		o.Rollback()
		return err
	}

	if languageNum == nil {
		*languageNum = 0
	}

	fmt.Println("languageNum: ", languageNum)
	//新增
	var insertSql string = " insert into sdt_bdi_rule_language ( " +
		"	bdi_id, " +
		//"	bdi_business_data_id, " +
		"	language_name, " +
		"	language_num, " +
		"	default_value, " +
		"	remarks, " +
		"	create_user_id, " +
		"	create_time ) " +
		"values (?, ?, ?, ?, ?, ?, ?) "

	_, err = o.Raw(insertSql, this.BdiId, this.LanguageName, *languageNum+1, this.DefaultValue, this.Remarks, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//更新
func (this *SdtBdiRuleLanguage) Update() error {
	o := orm.NewOrm()
	o.Begin()

	//var updateSql string = "update sdt_bdi_rule_set " +
	//"set rule_set_name = ?, " +
	//" rule_set_level = ?, " +
	//" remarks = ?, " +
	////" modify_user_id = ?, " + //修改人暂时不填
	//" modify_time = ? " +
	//"where bdi_rule_set_id = ?"
	//
	//_, err := o.Raw(updateSql, this.RuleName, this.BdiRuleSetId, this.Remarks, time.Now(), this.BdiRuleId).Exec()
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	o.Rollback()
	//	return err
	//}

	o.Commit()
	return nil
}
