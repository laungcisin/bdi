package controllers

import "bdi/models"

type AdtBdiAdmTreeController struct {
	BaseController
}

//获取行政区域树
func (this *AdtBdiAdmTreeController) Index() {
	tree := this.GetTree()
	this.Data["json"] = tree
	this.ServeJSON()
	return
}

func (this *AdtBdiAdmTreeController) GetTree() *[]models.Tree {
	id := this.GetString("id")

	var pid string
	if len(id) <= 0 {
		pid = "0"
	} else {
		pid = id
	}

	tree := models.GetSubAdtBdiAdmNode(pid)

	return tree
}
