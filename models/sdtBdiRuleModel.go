package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type SdtBdiRule struct {
	BdiRuleId           int       //主键
	BdiRuleSetId        int       //规则集主键
	RuleName string //规则名称

	BdiId               int       //指标项Id
	Condition           string
	ConditionExpression string
	Operation           string
	OperationExpression string
	Remarks             string    //指标集备注

	CreateUserId        int       //创建人ID
	CreateTime          time.Time //创建时间
	ModifyUserId        int       //修改人ID
	ModifyTime          time.Time //修改时间

				      //以下字段为datagrid展示
	BdiName             string    //指标项名称
}

func (u *SdtBdiRule) TableName() string {
	return "sdt_bdi_rule"
}

/**
获取所有指标。
*/
func (this *SdtBdiRule) GetAllSdtBdiRuleByRuleSetId(rows int, page int) ([]SdtBdiRule, int, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	sdtBdiRuleSlice := make([]SdtBdiRule, 0, 10)

	//查询的字段顺序最好和model的字段对应，方便解析并赋值。
	var querySql =
	" select r.*, s.rule_set_name " +
	" from " +
	"	sdt_bdi_rule r, " +
	"	sdt_bdi_rule_set s " +
	" where " +
	"	r.bdi_rule_set_id = ? " +
	" and r.bdi_rule_set_id = s.bdi_rule_set_id limit ?, ?"

	_, err := o.Raw(querySql, this.BdiRuleSetId, (page - 1) * rows, page * rows).QueryRows(&sdtBdiRuleSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	var countSql =
		" select count(*) as counts from " +
		"	sdt_bdi_rule r, " +
		"	sdt_bdi_rule_set s " +
		" where " +
		"	r.bdi_rule_set_id = ? " +
		" and r.bdi_rule_set_id = s.bdi_rule_set_id "

	var num *int = new(int)
	err = o.Raw(countSql, this.BdiRuleSetId).QueryRow(num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return sdtBdiRuleSlice, *num, nil
}


//新增
func (this *SdtBdiRule) Add() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	//新增前获取最大RuleSetLevel
	var querySql string = " select max(rule_set_level) as rule_set_level from sdt_bdi_rule_set where bdi_id = ? "
	maxLevel := new(int)
	err = o.Raw(querySql, this.BdiId).QueryRow(maxLevel)

	if err != nil {
		o.Rollback()
		return err
	}

	//新增
	var insertSql string = "insert into sdt_bdi_rule_set ( " +
	"	bdi_id, " +
	"	rule_set_name, " +
	"	rule_set_level, " +
	"	remarks, " +
	"	create_user_id, " +
	"	create_time " +
	") " +
	"values (?, ?, ?, ?, ?, ?)"

	_, err = o.Raw(insertSql, this.BdiId, this.RuleName, *maxLevel + 1, this.Remarks, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//更新
func (this *SdtBdiRule) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql string = "update sdt_bdi_rule_set " +
	"set rule_set_name = ?, " +
	" rule_set_level = ?, " +
	" remarks = ?, " +
	//" modify_user_id = ?, " + //修改人暂时不填
	" modify_time = ? " +
	"where bdi_rule_set_id = ?"

	_, err := o.Raw(updateSql, this.RuleName, this.BdiRuleSetId, this.Remarks, time.Now(), this.BdiRuleId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//行上移
func (this *SdtBdiRule) RowMoveUp() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	//先查询出rule_set_level 加 1 的行记录id
	var incLevelRuleSetIdSql string = " select bdi_rule_set_id from sdt_bdi_rule_set where bdi_id = ? and rule_set_level = ? limit 1 "
	incLevelRuleSetId := new(int)

	err = o.Raw(incLevelRuleSetIdSql, this.BdiId, this.BdiRuleSetId - 1).QueryRow(incLevelRuleSetId)
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	var updateSql string =
	" update sdt_bdi_rule_set " +
	" set rule_set_level = ?, " +
	//" modify_user_id = ?, " + //修改人暂时不填
	" modify_time = ? " +
	" where bdi_rule_set_id = ?"

	//1. 更新下移行记录的 BdiRuleSetId
	_, err = o.Raw(updateSql, this.BdiRuleSetId, time.Now(), incLevelRuleSetId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	//2. 更新上移行记录的 BdiRuleSetId
	_, err = o.Raw(updateSql, this.BdiRuleSetId - 1, time.Now(), this.BdiRuleId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//行下移
func (this *SdtBdiRule) RowMoveDown() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	//先查询出rule_set_level 减 1 的行记录id
	var decLevelRuleSetIdSql string = " select bdi_rule_set_id from sdt_bdi_rule_set where bdi_id = ? and rule_set_level = ? limit 1 "
	decLevelRuleSetId := new(int)

	err = o.Raw(decLevelRuleSetIdSql, this.BdiId, this.BdiRuleSetId + 1).QueryRow(decLevelRuleSetId)
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	var updateSql string =
	" update sdt_bdi_rule_set " +
	" set rule_set_level = ?, " +
	//" modify_user_id = ?, " + //修改人暂时不填
	" modify_time = ? " +
	" where bdi_rule_set_id = ?"

	//1. 更新下移行记录的 BdiRuleSetId
	_, err = o.Raw(updateSql, this.BdiRuleSetId, time.Now(), decLevelRuleSetId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	//2. 更新上移行记录的 BdiRuleSetId
	_, err = o.Raw(updateSql, this.BdiRuleSetId + 1, time.Now(), this.BdiRuleId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//行删除
func (this *SdtBdiRule) Delete() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	var updateSql string =
	" update sdt_bdi_rule_set " +
	" set rule_set_level = rule_set_level - 1, " +
	//" modify_user_id = ?, " + //修改人暂时不填
	" modify_time = ? " +
	" where bdi_id = ? and rule_set_level > ? "

	//1. BdiRuleSetId > 当前行，BdiRuleSetId - 1
	_, err = o.Raw(updateSql, time.Now(), this.BdiId, this.BdiRuleSetId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	//删除当前行
	var deleteSql string = " delete from sdt_bdi_rule_set where bdi_rule_set_id = ?  "
	_, err = o.Raw(deleteSql, this.BdiRuleId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}