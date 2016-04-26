package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"strings"
	"time"
)

type SdtBdiCfgKpi struct {
	KpiId   int    `form:"kpiId"`   //主键
	KpiName string `form:"kpiName"` //名称
	Remarks string `form:"remarks"` //备注
	BdiId   int    `form:"bdiId"`   //bdiId

	CreateUserId int       //创建人ID
	CreateTime   time.Time //创建时间
	ModifyUserId int       //修改人ID
	ModifyTime   time.Time //修改时间

	BdiCfgDimIds string `form:"bdiCfgDimIds"`

	//以下字段为datagrid展示
	BdiBaseName   string //所属标类库名称
	BdiDomainName string //所属标类域名称
	BdiSetName    string //所属指标集名称
	BdiName       string //所属指标项名称
}

func (u *SdtBdiCfgKpi) TableName() string {
	return "sdt_bdi_cfg_kpi"
}

/**
获取所有指标。
*/
func (this *SdtBdiCfgKpi) GetAllSdtBdiCfgKpi(rows int, page int) ([]SdtBdiCfgKpi, int, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	o.Begin()
	var sdtBdiCfgKpiSlice []SdtBdiCfgKpi

	var querySql = " select " +
		"	k.kpi_id, " +
		"	k.kpi_name, " +
		"	k.remarks, " +
		"	b.bdi_base_name, " +
		"	d.bdi_domain_name, " +
		"	s.bdi_set_name, " +
		"	bdi.bdi_name " +
		" from " +
		"	sdt_bdi_cfg_kpi k, " +
		"	sdt_bdi bdi, " +
		"	sdt_bdi_set_rel_bdi srb, " +
		"	sdt_bdi_set s, " +
		"	sdt_bdi_domain_set bds, " +
		"	sdt_bdi_domain d, " +
		"	sdt_bdi_base b " +
		" where " +
		" k.bdi_id = bdi.bdi_id " +
		" and bdi.bdi_id = srb.bdi_id " +
		" and srb.bdi_set_id = s.bdi_set_id " +
		" and s.bdi_set_id = bds.bdi_set_id " +
		" and bds.bdi_domain_id = d.bdi_domain_id " +
		" and d.bdi_base_id = b.bdi_base_id limit ?, ?; "

	num, err := o.Raw(querySql, (page-1)*rows, page*rows).QueryRows(&sdtBdiCfgKpiSlice)
	if err != nil {
		fmt.Println("err: ", err)
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	var countSql = " select count(*) as counts from " +
		"	sdt_bdi_cfg_kpi k, " +
		"	sdt_bdi bdi, " +
		"	sdt_bdi_set_rel_bdi srb, " +
		"	sdt_bdi_set s, " +
		"	sdt_bdi_domain_set bds, " +
		"	sdt_bdi_domain d, " +
		"	sdt_bdi_base b " +
		" where " +
		" k.bdi_id = bdi.bdi_id " +
		" and bdi.bdi_id = srb.bdi_id " +
		" and srb.bdi_set_id = s.bdi_set_id " +
		" and s.bdi_set_id = bds.bdi_set_id " +
		" and bds.bdi_domain_id = d.bdi_domain_id " +
		" and d.bdi_base_id = b.bdi_base_id "

	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		fmt.Println("err: ", err)
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	o.Commit()
	return sdtBdiCfgKpiSlice, int(num), nil
}

/**
获取所有指标。
*/
func (this *SdtBdiCfgKpi) GetSdtBdiCfgKpiById() error {
	// var o orm.Ormer
	// o = orm.NewOrm()
	// //查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	// var querySql = "select " +
	// 	"	sb.bdi_id, " +
	// 	"	sb.bdi_name, " +
	// 	"	sb.bdi_type_id, " +
	// 	"	sb.bdi_secrecy_id, " +
	// 	"	sb.remarks, " +
	// 	"	sbt.bdi_type_name, " +
	// 	"	sbsec.bdi_secrecy_name, " +
	// 	"	group_concat(sbset.bdi_set_id) as bdi_set_ids, " +
	// 	"	group_concat(sbset.bdi_set_name) as bdi_set_names " +
	// 	"from " +
	// 	"	sdt_bdi sb, " +
	// 	"	sdt_bdi_set_rel_bdi rel, " +
	// 	"	sdt_bdi_set sbset, " +
	// 	"	sdt_bdi_type sbt, " +
	// 	"	sdt_bdi_secrecy sbsec " +
	// 	"where " +
	// 	"sb.bdi_id = ? " +
	// 	"and sb.bdi_id = rel.bdi_id " +
	// 	"and rel.bdi_set_id = sbset.bdi_set_id " +
	// 	"and sb.bdi_type_id = sbt.bdi_type_id " +
	// 	"and sb.bdi_secrecy_id = sbsec.bdi_secrecy_id "

	// err := o.Raw(querySql, this.BdiId).QueryRow(this)
	// if err != nil {
	// 	log.Fatal("查询表：sdt_bdi_set出错！")
	// 	return err
	// }

	return nil
}

func (this *SdtBdiCfgKpi) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertCfgKpiSql = " insert into sdt_bdi_cfg_kpi ( " +
		"	kpi_name, " +
		"	remarks, " +
		"   bdi_id, " +
		"   create_user_id, " +
		"   create_time ) " +
		" values (?, ?, ?, ?, ?) "

	res, err := o.Raw(insertCfgKpiSql, this.KpiName, this.Remarks, this.BdiId, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	kpiId, err := res.LastInsertId()

	if err != nil {
		o.Rollback()
		return err
	}

	var insertCfgKpiRelSql = " insert into sdt_bdi_cfg_dim (kpi_id, bdi_stat_dim_id) values(?, ?) "
	bdiCfgDimIds := strings.Split(this.BdiCfgDimIds, ",")
	for _, v := range bdiCfgDimIds {
		bdiCfgDimId, err := strconv.Atoi(v)
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.Raw(insertCfgKpiRelSql, kpiId, bdiCfgDimId).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	o.Commit()
	return nil
}

func (this *SdtBdiCfgKpi) Update() error {
	// o := orm.NewOrm()
	// o.Begin()

	// var deleteSdtBdiSetRelBdiSql = " delete from sdt_bdi_set_rel_bdi where bdi_id = ? "
	// var updateSdtBdiSql = "update sdt_bdi " +
	// 	"set  " +
	// 	" bdi_name = ?, " +
	// 	" bdi_type_id = ?, " +
	// 	" bdi_secrecy_id = ?, " +
	// 	" remarks = ?, " +
	// 	//" modify_user_id = ?, " + //暂时没有修改人
	// 	" modify_time = ? " +
	// 	"where bdi_id = ?"
	// var insertSdtBdiSetRelBdiSql = " insert into sdt_bdi_set_rel_bdi(bdi_set_id, bdi_id) values (?, ?) "

	// _, err := o.Raw(updateSdtBdiSql, this.BdiName, this.BdiTypeId, this.BdiSecrecyId, this.Remarks, time.Now(), this.BdiId).Exec()
	// if err != nil {
	// 	o.Rollback()
	// 	return err
	// }

	// _, err = o.Raw(deleteSdtBdiSetRelBdiSql, this.BdiId).Exec()
	// if err != nil {
	// 	o.Rollback()
	// 	return err
	// }

	// bdiSetIds := strings.Split(this.BdiSetIds, ",")

	// for _, v := range bdiSetIds {
	// 	bdiSetId, err := strconv.Atoi(v)
	// 	if err != nil {
	// 		o.Rollback()
	// 		return err
	// 	}
	// 	_, err = o.Raw(insertSdtBdiSetRelBdiSql, bdiSetId, this.BdiId).Exec()
	// 	if err != nil {
	// 		o.Rollback()
	// 		return err
	// 	}
	// }

	// o.Commit()
	return nil
}
