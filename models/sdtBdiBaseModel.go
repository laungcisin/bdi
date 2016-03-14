package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type SdtBdiBase struct {
	BdiBaseId    int       `orm:"pk;auto";form:"bdiBaseId";json:"BdiBaseId"` //主键，自动增长
	BdiBaseName  string    `form:"bdiBaseName";json:"BdiBaseName"`           //指标库名称
	BdiBaseYear  string    `form:"bdiBaseYear";json:"BdiBaseYear"`           //统计年
	AdmCode      string    `form:"admCode";json:"AdmCode"`                   //统计行政区域代码
	AdtFlag      string    `form:"adtFlag";json:"AdtFlag"`                   //是否使用: 0 - 停用， 1 - 使用
	ManagerId    int       `form:"managerId";json:"ManagerId"`               //管理员ID
	CreateUserId int       `form:"createUserId";json:"CreateUserId"`         //创建人ID
	CreateTime   time.Time `form:"createTime";json:"CreateTime"`             //创建时间
	ModifyUserId int       `form:"modifyUserId";json:"ModifyUserId"`         //修改人ID
	ModifyTime   time.Time `form:"modifyTime";json:"ModifyTime"`             //修改时间
	Remarks      string    `form:"remarks";json:"Remarks"`                   //指标库备注
}

func (u *SdtBdiBase) TableName() string {
	return "sdt_bdi_base"
}

/**
获取所有指标库-用于下拉列表。
 */
func GetAllSdtBdiBaseForList() ([]orm.Params, error) {
	var o orm.Ormer
	o = orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(" select bdi_base_id as Id, bdi_base_name as Text from sdt_bdi_base ").Values(&maps)

	if err != nil {
		return nil, err
	}

	return maps, nil
}
