package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type SdtBdiBusiConfig struct {
	Id                int    `form:"id"`                //主键
	BusiId            int    `form:"busiId"`            //业务表ID
	Name              string `form:"name"`              //配置业务表名
	CnName            string `form:"cnName"`            //配置业务表中文名
	IsProcess         int    `form:"isProcess"`         //是否经过处理
	SelectStar        string `form:"selectStar"`        //选择字段
	ProcessColumn     string `form:"processColumn"`     //处理字段
	AsName            string `form:"asName"`            //别名
	ProcessDataType   string `form:"processDataType"`   //
	ProcessDataLength string `form:"processDataLength"` //
	ProcessType       string `form:"processType"`       //处理类型
	Operators         string `form:"operators"`         //操作符
	Params            string `form:"params"`            //参数

	UserCode   string    `form:"userCode"`   //创建人ID
	CreateTime time.Time `form:"createTime"` //创建时间
	EditTime   time.Time `form:"editTime"`   //修改时间

	//以下字段为datagrid展示
	BdiId    int `form:"bdiId"` //
	BusiName string
}

func (u *SdtBdiBusiConfig) TableName() string {
	return "sdt_bdi_busi_config"
}

/**
获取所有指标。
*/
func (this *SdtBdiBusiConfig) GetAllSdtBdiBusiConfig(rows int, page int) ([]SdtBdiBusiConfig, int, error) {
	o := orm.NewOrm()
	var sdtBdiBusiConfigSlice []SdtBdiBusiConfig = make([]SdtBdiBusiConfig, 0)

	var querySql = " select " +
		"	config.*,  " +
		"	busi.name as busi_name, " +
		"	bdi.bdi_name " +
		" from " +
		"	sdt_bdi_busi_config config, " +
		"	sdt_bdi_busi busi, " +
		"	sdt_bdi bdi " +
		" where " +
		"	bdi.id = ? " +
		" and config.busi_id = busi.id " +
		" and busi.bdi_id = bdi.id " +
		"  limit ?, ? "

	_, err := o.Raw(querySql, this.BdiId, (page-1)*rows, page*rows).QueryRows(&sdtBdiBusiConfigSlice)
	if err != nil {
		fmt.Println("查询表：" + this.TableName() + "出错！")
		fmt.Println(err)
		return nil, 0, err
	}

	num := new(int)
	var countSql = " select count(*) as counts " +
		" from " +
		"	sdt_bdi_busi_config config, " +
		"	sdt_bdi_busi busi, " +
		"	sdt_bdi bdi " +
		" where " +
		"	bdi.id = ? " +
		" and config.busi_id = busi.id " +
		" and busi.bdi_id = bdi.id "

	err = o.Raw(countSql, this.BdiId).QueryRow(&num)
	if err != nil {
		fmt.Println("查询表：" + this.TableName() + "出错！")
		fmt.Println(err)
		return nil, 0, err
	}

	return sdtBdiBusiConfigSlice, *num, nil
}

/**
 */
func (this *SdtBdiBusiConfig) GetSdtBdiBusiConfigById() error {
	o := orm.NewOrm()
	var querySql = " select * from sdt_bdi_busi_config where id = ? "

	err := o.Raw(querySql, this.Id).QueryRow(this)
	if err != nil {
		fmt.Println("查询表：" + this.TableName() + "出错！")
		return err
	}

	return nil
}

func (this *SdtBdiBusiConfig) Update() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	var updateString = " update sdt_bdi_busi_config set   " +
		" name = ?, " +
		" cn_name = ?, " +
		" is_process = 1, " +
		" process_column = ?, " +
		" process_data_type = ?, " +
		" process_data_length = ?, " +
		" operators = ?, " +
		" select_star = ?, " +
		" as_name = ?, " +
		" process_type = ?, " +
		" params = ?, " +
		" user_code = ?, " +
		" edit_time = ? " +
		" where id = ? "
	_, err = o.Raw(updateString, this.Name, this.CnName,this.ProcessColumn, this.ProcessDataType,
			this.ProcessDataLength, this.Operators, this.SelectStar, this.AsName, this.ProcessType, this.Params, 0,
		time.Now(), this.Id).Exec()
	if err != nil {
		fmt.Println(err)
		fmt.Println("更新数据出错！")
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

func (this *SdtBdiBusiConfig) Synchronize() error {
	var err error
	o := orm.NewOrm()
	o.Begin()

	var insertString = " insert into sdt_bdi_busi_fields ( " +
		"	busi_id, " +
		"       bdi_id, " +
		"	name, " +
		"	sequence, " +
		"       is_process, " +
		"	comment, " +
		"	data_type, " +
		"	data_length, " +
		"	create_time " +
		" ) " +
		" select t.* from ( " +
		" select cfg.busi_id as busi_id, " +
		"       busi.bdi_id, " +
		"	cfg.as_name as name, " +
		"	1 as sequence, " +
		"	1 as is_process, " +
		"	cfg.cn_name as comment, " +
		"	cfg.process_data_type as data_type, " +
		"	cfg.process_data_length as data_length, " +
		"	now() as create_time " +
		"	from " +
		"		sdt_bdi_busi_config cfg, " +
		"		sdt_bdi_busi busi " +
		"	where " +
		"		busi.bdi_id = ? " +
		"	and busi.id = cfg.busi_id " +
		"	group by cfg.as_name " +
		") as t where not exists ( select f.busi_id from sdt_bdi_busi_fields f where f.busi_id = t.busi_id and f.name = t.name) "
	_, err = o.Raw(insertString, this.BdiId).Exec()
	if err != nil {
		fmt.Println(err)
		fmt.Println("更新数据出错！")
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}
