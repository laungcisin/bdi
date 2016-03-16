package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type SdtBdiRuleSet struct {
	BdiRuleSetId int    //主键
	BdiId        int    //指标项Id
	RuleSetName  string //规则集名称
	RuleSetLevel int    //规则集执行顺序
	Remarks      string //指标集备注

	CreateUserId int       //创建人ID
	CreateTime   time.Time //创建时间
	ModifyUserId int       //修改人ID
	ModifyTime   time.Time //修改时间

	//以下字段为datagrid展示
	BdiName string //指标项名称

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

	num, err := o.Raw(querySql, this.BdiId, (page-1)*rows, page*rows).QueryRows(&sdtBdiRuleSetSlice)
	fmt.Println("sdtBdiRuleSetSlice: ", sdtBdiRuleSetSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	var countSql = " select count(*) as counts from " + this.TableName() + " where bdi_id = ? "
	err = o.Raw(countSql, this.BdiId).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return sdtBdiRuleSetSlice, int(num), nil
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
