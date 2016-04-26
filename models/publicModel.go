package models
import (
	"github.com/astaxie/beego/orm"
)

/**
获取所有指标库-用于下拉列表。
*/
func GetAllSdtBdiBaseForList() ([]orm.Params, error) {
	o := orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select id as Id, bdi_base_name as Text from sdt_bdi_base ").Values(&maps)

	if err != nil {
		return nil, err
	}

	return maps, nil
}

/**
获取所有指标域-用于下拉列表。
*/
func GetAllSdtBdiDomainForList() ([]orm.Params, error) {
	o := orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select id as Id, bdi_domain_name as Text from sdt_bdi_domain ").Values(&maps)

	if err != nil {
		return nil, err
	}

	return maps, nil
}

