package models

import (
//	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type SdtBdiDomain struct {
	BdiDomainId   int       `orm:"pk;auto;column(bdi_domain_id)";form:"-";json:"BdiDomainId"`                       //主键，自动增长
	BdiBaseId     int       `orm:"column(bdi_base_id);null";form:"bdiBaseId";json:"BdiBaseId"`                      //指标库Id
	BdiBaseName   string    `orm:"column(bdi_base_name);null";form:"bdiBaseName";json:"BdiBaseName"`                                         //指标库名称
	BdiDomainName string    `orm:"column(bdi_domain_name);null";form:"bdiDomainName";json:"BdiDomainName"`          //指标域名称
	AdtFlag       string    `orm:"column(adt_flag);null";form:"adtFlag";json:"AdtFlag"`                             //是否使用: 0 - 停用， 1 - 使用
	Remarks       string    `orm:"column(remarks);null";form:"remarks";json:"Remarks"`                              //指标库备注
	CreateUserId  int       `orm:"column(create_user_id);null";form:"-";json:"CreateUserId"`                        //创建人ID
	CreateTime    time.Time `orm:"column(create_time);null;auto_now_add;type(datetime)";form:"-";json:"CreateTime"` //创建时间
	ModifyUserId  int       `orm:"column(modify_user_id);null";form:"-";json:"ModifyUserId"`                        //修改人ID
	ModifyTime    time.Time `orm:"column(modify_time);null;auto_now;type(datetime)";form:"-";json:"ModifyTime"`     //修改时间
}

func (u *SdtBdiDomain) TableName() string {
	return "sdt_bdi_domain"
}

/**
获取所有指标域。
 */
func GetAllSdtBdiDomain(rows int, page int) ([]SdtBdiDomain, int, error) {
	var o orm.Ormer
	o = orm.NewOrm()
	var sdtBdiDomainList []SdtBdiDomain


	//查询的字段顺序最好和model的字段对应，这样才方便解析并赋值。
	var querySql = "SELECT " +
	"	sdt_bdi_domain.bdi_domain_id, " +
	"	sdt_bdi_domain.bdi_base_id, " +
	"	sdt_bdi_base.bdi_base_name, " +
	"	sdt_bdi_domain.bdi_domain_name, " +
	"	sdt_bdi_domain.adt_flag, " +
	"	sdt_bdi_domain.remarks, " +
	"	sdt_bdi_domain.create_user_id, " +
	"	sdt_bdi_domain.create_time, " +
	"	sdt_bdi_domain.modify_user_id, " +
	"	sdt_bdi_domain.modify_time " +
	"FROM 	sdt_bdi_domain " +
	"LEFT JOIN sdt_bdi_base ON sdt_bdi_domain.bdi_base_id = sdt_bdi_base.bdi_base_id " + //关联sdt_bdi_base表
	"LIMIT ?, ?"

	num, err := o.Raw(querySql, (page - 1) * rows, page * rows).QueryRows(&sdtBdiDomainList)
	if err != nil {
		log.Fatal("查询表：sdt_bdi_base出错！")
		return nil, 0, err
	}

	var countSql = " select count(*) as counts from sdt_bdi_base "
	err = o.Raw(countSql).QueryRow(&num)
	if err != nil {
		log.Fatal("查询表：sdt_bdi_base出错！")
		return nil, 0, err
	}

	return sdtBdiDomainList, int(num), nil
}

func AddSdtBdiDomain(sdtBdiDomain *SdtBdiDomain) (error) {
	var o orm.Ormer
	o = orm.NewOrm()

	var insertSql = " insert into sdt_bdi_domain " +
	" (bdi_base_id , bdi_domain_name , adt_flag , remarks, create_time) values (?, ?, ?, ?, ?) "

	_, err := o.Raw(insertSql, sdtBdiDomain.BdiBaseId, sdtBdiDomain.BdiBaseName, sdtBdiDomain.AdtFlag, sdtBdiDomain.Remarks, time.Now()).Exec()
	if err != nil {
		return err
	}

	return nil
}

/**
获取所有指标域-用于下拉列表。
 */
func GetAllSdtBdiDomainForList() ([]orm.Params, error) {
	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select bdi_domain_id as Id, bdi_domain_name as Text from sdt_bdi_domain ").Values(&maps)

	if err != nil {
		return nil, err
	}

	return maps, nil
}
