package controllers

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