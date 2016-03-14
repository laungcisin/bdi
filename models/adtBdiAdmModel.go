package models

import (
	"github.com/astaxie/beego/orm"
	//	"fmt"
)

type AdtBdiAdm struct {
	AdmCode    string `orm:"pk";form:"admCode"` //主键
	ParentCode string `form:"parentCode"`       //父Id
	AdmName    string `form:"admName"`          //区域名称
}

type Tree struct {
	Id       string `json:"id"`
	Pid      string `json:"pid"`
	Text     string `json:"text"`
	IconCls  string `json:"iconCls"`
	Checked  string `json:"checked"`
	State    string `json:"state"`
	Children []Tree `json:"children"`
}

func (u *AdtBdiAdm) TableName() string {
	return "adt_bdi_adm"
}

func GetAdtBdiAdmNode(parentCode string) ([]orm.Params, error) {
	o := orm.NewOrm()
	adtBdiAdm := new(AdtBdiAdm)
	var adtBdiAdmSlice []orm.Params
	_, err := o.QueryTable(adtBdiAdm.TableName()).Filter("ParentCode", parentCode).Values(&adtBdiAdmSlice)
	if err != nil {
		return adtBdiAdmSlice, err
	}
	return adtBdiAdmSlice, nil
}

/**
获取一个节点下的所有子节点。
*/
func GetAdtBdiAdmTree(parentCode string, treeSlice *[]Tree) *[]Tree {
	//获取一个节点下所有的子节点
	o := orm.NewOrm()
	adtBdiAdm := new(AdtBdiAdm)
	var adtBdiAdmSlice []orm.Params
	o.QueryTable(adtBdiAdm.TableName()).Filter("ParentCode", parentCode).Values(&adtBdiAdmSlice)

	//	fmt.Println("AdtBdiAdm: ", adtBdiAdmSlice)

	for _, v := range adtBdiAdmSlice {
		//		fmt.Print("k: ", k)
		//		fmt.Print("      ")
		//		fmt.Print("v: ", v)
		//		fmt.Println()

		tree := new(Tree)
		tree.Id = v["AdmCode"].(string)
		tree.Text = v["AdmName"].(string)
		tree.Pid = v["ParentCode"].(string)
		tree.State = "closed"

		//获取子节点数量
		num, _ := GetSubAdtBdiAdmTreeCount(tree.Id)
		//不包含子节点
		if num < 1 {
		} else {
			newTreeDataSlice := []Tree{}
			tempTree := GetAdtBdiAdmTree(tree.Id, &newTreeDataSlice)
			tree.Children = *tempTree
		}

		*treeSlice = append(*treeSlice, *tree)
	}

	//	fmt.Println("treeSlice:" , treeSlice)
	return treeSlice
}

/**
节点下是否有子节点
*/
func GetSubAdtBdiAdmTreeCount(admCode string) (int64, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable(new(AdtBdiAdm).TableName()).Filter("ParentCode", admCode).Count()
	if err != nil {
		return -1, err
	}

	return num, nil
}

/**
获取一个节点下的所有子节点。
*/
func GetSubAdtBdiAdmNode(parentCode string) *[]Tree {
	//获取一个节点下所有的子节点
	o := orm.NewOrm()
	adtBdiAdm := new(AdtBdiAdm)
	var adtBdiAdmSlice []orm.Params
	_, err := o.QueryTable(adtBdiAdm.TableName()).Filter("ParentCode", parentCode).Values(&adtBdiAdmSlice)

	//	fmt.Println("AdtBdiAdm: ", adtBdiAdmSlice)
	if err != nil {
		return nil
	}

	newTreeDataSlice := []Tree{}
	for _, v := range adtBdiAdmSlice {
		//		fmt.Print("k: ", k)
		//		fmt.Print("      ")
		//		fmt.Print("v: ", v)
		//		fmt.Println()

		treeNode := new(Tree)
		treeNode.Id = v["AdmCode"].(string)
		treeNode.Text = v["AdmName"].(string)
		treeNode.Pid = v["ParentCode"].(string)

		//获取子节点数量
		num, _ := GetSubAdtBdiAdmTreeCount(treeNode.Id)
		//不包含子节点
		if num >= 1 {
			treeNode.State = "closed"
		}
		newTreeDataSlice = append(newTreeDataSlice, *treeNode)
	}

	//	fmt.Println("treeSlice:" , treeSlice)
	return &newTreeDataSlice
}
