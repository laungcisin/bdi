package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type SdtBdiRuleSet struct {
	BdiRuleSetId int       //主键
	BdiId        int       //指标项Id
	RuleSetName  string    //规则集名称
	RuleSetLevel int       //规则集执行顺序
	Remarks      string    //指标集备注

	CreateUserId int       //创建人ID
	CreateTime   time.Time //创建时间
	ModifyUserId int       //修改人ID
	ModifyTime   time.Time //修改时间

			       //以下字段为datagrid展示
	BdiName      string    //指标项名称
}

func (u *SdtBdiRuleSet) TableName() string {
	return "sdt_bdi_rule_set"
}

/**
获取所有指标。
*/
func (this *SdtBdiRuleSet) GetAllSdtBdiRuleSetByBdiId(rows int, page int) ([]SdtBdiRuleSet, int, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	sdtBdiRuleSetSlice := make([]SdtBdiRuleSet, 0, 0)

	//查询的字段顺序最好和model的字段对应，方便解析并赋值。
	var querySql = " select rs.*, bdi.bdi_name " +
	" from " +
	"	sdt_bdi_rule_set rs, " +
	"	sdt_bdi bdi " +
	" where rs.bdi_id = ? and rs.bdi_id = bdi.bdi_id order by rs.rule_set_level limit ?, ? "

	_, err := o.Raw(querySql, this.BdiId, (page - 1) * rows, page * rows).QueryRows(&sdtBdiRuleSetSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	var countSql = " select count(*) as counts from " + this.TableName() + " where bdi_id = ? "

	var num *int = new(int)
	err = o.Raw(countSql, this.BdiId).QueryRow(num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return sdtBdiRuleSetSlice, *num, nil
}

//更新规则集表-新增-删除-修改
func (this *SdtBdiRuleSet) UpdateStatus(insertedRuleSetSlice []SdtBdiRuleSet, deletedRuleSetSlice []SdtBdiRuleSet, updatedRuleSetSlice []SdtBdiRuleSet) error {
	o := orm.NewOrm()
	o.Begin()

	//新增
	if len(insertedRuleSetSlice) > 0 {
		var insertSql string = "insert into sdt_bdi_rule_set ( " +
		"	bdi_id, " +
		"	rule_set_name, " +
		"	rule_set_level, " +
		"	remarks, " +
		"	create_user_id, " +
		"	create_time " +
		") " +
		"values (?, ?, ?, ?, ?, ?)"

		for _, value := range insertedRuleSetSlice {
			_, err := o.Raw(insertSql, value.BdiId, value.RuleSetName, value.RuleSetLevel, value.Remarks, 0, time.Now()).Exec()
			if err != nil {
				o.Rollback()
				return err
			}
		}
	}

	//删除
	if len(deletedRuleSetSlice) > 0 {
		var deleteSql string = " delete from sdt_bdi_rule_set where bdi_rule_set_id = ? "

		for _, value := range deletedRuleSetSlice {
			_, err := o.Raw(deleteSql, value.BdiRuleSetId).Exec()
			if err != nil {
				fmt.Println("err: ", err)
				o.Rollback()
				return err
			}
		}
	}

	//修改
	if len(updatedRuleSetSlice) > 0 {
		var updateSql string = "update sdt_bdi_rule_set " +
		"set rule_set_name = ?, " +
		" rule_set_level = ?, " +
		" remarks = ?, " +
		//" modify_user_id = ?, " + //修改人暂时不填
		" modify_time = ? " +
		"where bdi_rule_set_id = ?"

		for _, value := range updatedRuleSetSlice {
			_, err := o.Raw(updateSql, value.RuleSetName, value.RuleSetLevel, value.Remarks, time.Now(), value.BdiRuleSetId).Exec()
			if err != nil {
				fmt.Println("err: ", err)
				o.Rollback()
				return err
			}
		}
	}

	o.Commit()
	return nil
}

//新增
func (this *SdtBdiRuleSet) Add() error {
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

	_, err = o.Raw(insertSql, this.BdiId, this.RuleSetName, *maxLevel + 1, this.Remarks, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//更新
func (this *SdtBdiRuleSet) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql string = "update sdt_bdi_rule_set " +
	"set rule_set_name = ?, " +
	" rule_set_level = ?, " +
	" remarks = ?, " +
	//" modify_user_id = ?, " + //修改人暂时不填
	" modify_time = ? " +
	"where bdi_rule_set_id = ?"

	_, err := o.Raw(updateSql, this.RuleSetName, this.RuleSetLevel, this.Remarks, time.Now(), this.BdiRuleSetId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//行上移
func (this *SdtBdiRuleSet) RowMoveUp() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	//先查询出rule_set_level 加 1 的行记录id
	var incLevelRuleSetIdSql string = " select bdi_rule_set_id from sdt_bdi_rule_set where bdi_id = ? and rule_set_level = ? limit 1 "
	incLevelRuleSetId := new(int)

	err = o.Raw(incLevelRuleSetIdSql, this.BdiId, this.RuleSetLevel - 1).QueryRow(incLevelRuleSetId)
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

	//1. 更新下移行记录的 RuleSetLevel
	_, err = o.Raw(updateSql, this.RuleSetLevel, time.Now(), incLevelRuleSetId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	//2. 更新上移行记录的 RuleSetLevel
	_, err = o.Raw(updateSql, this.RuleSetLevel - 1, time.Now(), this.BdiRuleSetId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//行下移
func (this *SdtBdiRuleSet) RowMoveDown() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	//先查询出rule_set_level 减 1 的行记录id
	var decLevelRuleSetIdSql string = " select bdi_rule_set_id from sdt_bdi_rule_set where bdi_id = ? and rule_set_level = ? limit 1 "
	decLevelRuleSetId := new(int)

	err = o.Raw(decLevelRuleSetIdSql, this.BdiId, this.RuleSetLevel + 1).QueryRow(decLevelRuleSetId)
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

	//1. 更新下移行记录的 RuleSetLevel
	_, err = o.Raw(updateSql, this.RuleSetLevel, time.Now(), decLevelRuleSetId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	//2. 更新上移行记录的 RuleSetLevel
	_, err = o.Raw(updateSql, this.RuleSetLevel + 1, time.Now(), this.BdiRuleSetId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//行删除
func (this *SdtBdiRuleSet) Delete() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	var updateSql string =
		" update sdt_bdi_rule_set " +
		" set rule_set_level = rule_set_level - 1, " +
		//" modify_user_id = ?, " + //修改人暂时不填
		" modify_time = ? " +
		" where bdi_id = ? and rule_set_level > ? "

	//1. RuleSetLevel > 当前行，RuleSetLevel - 1
	_, err = o.Raw(updateSql, time.Now(), this.BdiId, this.RuleSetLevel).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	//删除当前行
	var deleteSql string = " delete from sdt_bdi_rule_set where bdi_rule_set_id = ?  "
	_, err = o.Raw(deleteSql, this.BdiRuleSetId).Exec()
	if err != nil {
		fmt.Println("err: ", err)
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}