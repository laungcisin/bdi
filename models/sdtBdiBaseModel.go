package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type SdtBdiBase struct {
	Id          int    `form:"id""`         //主键，自增长
	BdiBaseName string `form:"bdiBaseName"` //指标库名称
	BdiBaseYear string `form:"bdiBaseYear"` //统计年
	AdmCode     string `form:"admCode"`     //统计行政区域代码
	AdtFlag     byte   `form:"adtFlag"`     //是否使用: 1 - 启用， 0 - 停用
	Remarks     string `form:"remarks"`     //指标库备注

	UserCode   int       `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

}

func (this *SdtBdiBase) TableName() string {
	return "sdt_bdi_base"
}

/**
获取所有指标库--分页查询。
*/
func (this *SdtBdiBase) GetAllSdtBdiBase(rows int, page int) ([]SdtBdiBase, int, error) {
	o := orm.NewOrm()

	var sdtBdiBaseSlice []SdtBdiBase = make([]SdtBdiBase, 0)
	var querySql = " select * from " + this.TableName() + " limit ?, ?"

	_, err := o.Raw(querySql, (page-1)*rows, page*rows).QueryRows(&sdtBdiBaseSlice)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	num := new(int)
	var countSql = " select count(*) as counts from " + this.TableName()
	err = o.Raw(countSql).QueryRow(num)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return nil, 0, err
	}

	return sdtBdiBaseSlice, int(*num), nil
}

func (this *SdtBdiBase) GetSdtBdiBaseById() error {
	o := orm.NewOrm()
	var querySql = " select * from " + this.TableName() + " where id = ? "

	err := o.Raw(querySql, this.Id).QueryRow(this)
	if err != nil {
		log.Fatal("查询表：" + this.TableName() + "出错！")
		return err
	}

	return nil
}

func (this *SdtBdiBase) Add() error {
	o := orm.NewOrm()
	o.Begin()

	var insertSql = " insert into sdt_bdi_base ( " +
		"	bdi_base_name, " +
		"	bdi_base_year, " +
		"	adm_code, " +
		"	adt_flag, " +
		"	remarks, " +
		"	user_code, " +
		"	create_time) " +
		" values(?, ?, ?, ?, ?, ?, ?) "

	_, err := o.Raw(insertSql, this.BdiBaseName, this.BdiBaseYear, this.AdmCode, this.AdtFlag, this.Remarks, 0, time.Now()).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

func (this *SdtBdiBase) Update() error {
	o := orm.NewOrm()
	o.Begin()

	var updateSql = " update sdt_bdi_base " +
		" set bdi_base_name = ?, " +
		" bdi_base_year = ?, " +
		" adm_code = ?, " +
		" adt_flag = ?, " +
		" remarks = ?, " +
		" user_code = ?, " +
		" edit_time = ? " +
		" where id = ? "

	_, err := o.Raw(updateSql, this.BdiBaseName, this.BdiBaseYear, this.AdmCode, this.AdtFlag, this.Remarks, 0, time.Now(), this.Id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}
