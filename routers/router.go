package routers

import (
	"bdi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 路由设置
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")

	//区域树
	beego.Router("/bdi/adtBdiAdmTree", &controllers.AdtBdiAdmTreeController{}, "*:Index")

	//公共展示数据列表
	beego.Router("/sdt/sdtBdiBaseForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiBaseForList")       //列出所有的指标库-用于下拉列表
	beego.Router("/sdt/sdtBdiDomainForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiDomainForList")   //列出所有的指标域-用于下拉列表
	beego.Router("/sdt/sdtBdiSetForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiSetForList")         //列出所有的指标集-用于下拉列表
	beego.Router("/sdt/sdtBdiSecrecyForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiSecrecyForList") //列出所有的保密级别-用于下拉列表
	beego.Router("/sdt/sdtTypeSetForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiTypeForList")       //列出所有的指标类型-用于下拉列表

	//指标库
	beego.Router("/sdtBdiBase/index", &controllers.SdtBdiBaseController{}, "*:Index")               //列表首页
	beego.Router("/sdtBdiBase/all", &controllers.SdtBdiBaseController{}, "*:All")                   //列出所有数据
	beego.Router("/sdtBdiBase/content", &controllers.SdtBdiBaseController{}, "*:Content")           //编辑首页
	beego.Router("/sdtBdiBase/content/baseInfo", &controllers.SdtBdiBaseController{}, "*:BaseInfo") //获取基础信息
	beego.Router("/sdtBdiBase/update", &controllers.SdtBdiBaseController{}, "*:Update")             //更新信息
	beego.Router("/sdtBdiBase/add", &controllers.SdtBdiBaseController{}, "*:Add")                   //新增
	beego.Router("/sdtBdiBase/delete", &controllers.SdtBdiBaseController{}, "*:Delete")             //删除

	//指标域
	beego.Router("/sdtBdiDomain/index", &controllers.SdtBdiDomainController{}, "*:Index")   //列表首页
	beego.Router("/sdtBdiDomain/all", &controllers.SdtBdiDomainController{}, "*:All")       //列出所有数据
	beego.Router("/sdtBdiDomain/update", &controllers.SdtBdiDomainController{}, "*:Update") //更新
	beego.Router("/sdtBdiDomain/add", &controllers.SdtBdiDomainController{}, "*:Add")       //新增
	beego.Router("/sdtBdiDomain/delete", &controllers.SdtBdiDomainController{}, "*:Delete") //删除

	//指标集
	beego.Router("/sdtBdiSet/index", &controllers.SdtBdiSetController{}, "*:Index")           //列表首页
	beego.Router("/sdtBdiSet/all", &controllers.SdtBdiSetController{}, "*:All")               //列出所有数据
	beego.Router("/sdtBdiSet/updatePage", &controllers.SdtBdiSetController{}, "*:UpdatePage") //更新页面
	beego.Router("/sdtBdiSet/update", &controllers.SdtBdiSetController{}, "*:Update")         //更新内容
	beego.Router("/sdtBdiSet/addPage", &controllers.SdtBdiSetController{}, "*:AddPage")       //新增页面
	beego.Router("/sdtBdiSet/add", &controllers.SdtBdiSetController{}, "*:Add")               //新增内容
	beego.Router("/sdtBdiSet/delete", &controllers.SdtBdiSetController{}, "*:Delete")         //删除

	//指标项
	beego.Router("/sdtBdi/index", &controllers.SdtBdiController{}, "*:Index")           //列表首页
	beego.Router("/sdtBdi/all", &controllers.SdtBdiController{}, "*:All")               //列出所有数据
	beego.Router("/sdtBdi/updatePage", &controllers.SdtBdiController{}, "*:UpdatePage") //更新页面
	beego.Router("/sdtBdi/update", &controllers.SdtBdiController{}, "*:Update")         //更新内容
	beego.Router("/sdtBdi/addPage", &controllers.SdtBdiController{}, "*:AddPage")       //新增页面
	beego.Router("/sdtBdi/add", &controllers.SdtBdiController{}, "*:Add")               //新增内容

}
