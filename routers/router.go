package routers

import (
	"bdi/controllers"
	"bdi/models"
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
	beego.Router("/sdt/sdtBdiBaseForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiBaseForList")                         //列出所有的指标库-用于下拉列表
	beego.Router("/sdt/sdtBdiDomainForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiDomainForList")                     //列出所有的指标域-用于下拉列表
	beego.Router("/sdt/sdtBdiSetForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiSetForList")                           //列出所有的指标集-用于下拉列表
	beego.Router("/sdt/sdtBdiSecrecyForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiSecrecyForList")                   //列出所有的保密级别-用于下拉列表
	beego.Router("/sdt/sdtTypeSetForList", &controllers.SdtPublicDisplayDataController{}, "*:SdtBdiTypeForList")                         //列出所有的指标类型-用于下拉列表
	beego.Router("/sdt/datasourceTypeForList", &controllers.SdtPublicDisplayDataController{}, "*:DatasourceTypeForList")                 //列出所有的数据源类型-用于下拉列表
	beego.Router("/sdt/datasourceForList", &controllers.SdtPublicDisplayDataController{}, "*:DatasourceForList")                         //列出所有的数据源-用于下拉列表
	beego.Router("/sdt/processTypeForListByCondition", &controllers.SdtPublicDisplayDataController{}, "*:ProcessTypeForListByCondition") //列出所有的Stat-Dim-用于下拉列表
	beego.Router("/sdt/processTypeForListAll", &controllers.SdtPublicDisplayDataController{}, "*:ProcessTypeForListAll")                 //列出所有的Stat-Dim-用于下拉列表

	//指标库
	beego.Router("/sdtBdiBase/index", &controllers.SdtBdiBaseController{}, "*:Index")           //列表首页
	beego.Router("/sdtBdiBase/all", &controllers.SdtBdiBaseController{}, "*:All")               //列出所有数据
	beego.Router("/sdtBdiBase/addPage", &controllers.SdtBdiBaseController{}, "*:AddPage")       //新增页面
	beego.Router("/sdtBdiBase/add", &controllers.SdtBdiBaseController{}, "*:Add")               //新增
	beego.Router("/sdtBdiBase/updatePage", &controllers.SdtBdiBaseController{}, "*:UpdatePage") //更新页面
	beego.Router("/sdtBdiBase/update", &controllers.SdtBdiBaseController{}, "*:Update")         //更新信息
	beego.Router("/sdtBdiBase/delete", &controllers.SdtBdiBaseController{}, "*:Delete")         //删除

	//指标域
	beego.Router("/sdtBdiDomain/index", &controllers.SdtBdiDomainController{}, "*:Index")           //列表首页
	beego.Router("/sdtBdiDomain/all", &controllers.SdtBdiDomainController{}, "*:All")               //列出所有数据
	beego.Router("/sdtBdiDomain/updatePage", &controllers.SdtBdiDomainController{}, "*:UpdatePage") //更新页面
	beego.Router("/sdtBdiDomain/update", &controllers.SdtBdiDomainController{}, "*:Update")         //更新
	beego.Router("/sdtBdiDomain/addPage", &controllers.SdtBdiDomainController{}, "*:AddPage")       //新增页面
	beego.Router("/sdtBdiDomain/add", &controllers.SdtBdiDomainController{}, "*:Add")               //新增
	beego.Router("/sdtBdiDomain/delete", &controllers.SdtBdiDomainController{}, "*:Delete")         //删除

	//指标集
	beego.Router("/sdtBdiSet/index", &controllers.SdtBdiSetController{}, "*:Index")           //列表首页
	beego.Router("/sdtBdiSet/all", &controllers.SdtBdiSetController{}, "*:All")               //列出所有数据
	beego.Router("/sdtBdiSet/addPage", &controllers.SdtBdiSetController{}, "*:AddPage")       //新增页面
	beego.Router("/sdtBdiSet/add", &controllers.SdtBdiSetController{}, "*:Add")               //新增内容
	beego.Router("/sdtBdiSet/updatePage", &controllers.SdtBdiSetController{}, "*:UpdatePage") //更新页面
	beego.Router("/sdtBdiSet/update", &controllers.SdtBdiSetController{}, "*:Update")         //更新内容
	beego.Router("/sdtBdiSet/delete", &controllers.SdtBdiSetController{}, "*:Delete")         //删除

	//指标项
	beego.Router("/sdtBdi/index", &controllers.SdtBdiController{}, "*:Index")           //列表首页
	beego.Router("/sdtBdi/all", &controllers.SdtBdiController{}, "*:All")               //列出所有数据
	beego.Router("/sdtBdi/updatePage", &controllers.SdtBdiController{}, "*:UpdatePage") //更新页面
	beego.Router("/sdtBdi/update", &controllers.SdtBdiController{}, "*:Update")         //更新内容
	beego.Router("/sdtBdi/addPage", &controllers.SdtBdiController{}, "*:AddPage")       //新增页面
	beego.Router("/sdtBdi/add", &controllers.SdtBdiController{}, "*:Add")               //新增内容

	//指标类型
	beego.Router("/sdtBdiType/addPage", &controllers.SdtBdiTypeController{}, "*:AddPage")               //新增内容
	beego.Router("/sdtBdiType/add", &controllers.SdtBdiTypeController{}, "*:Add")               //新增内容


	//数据源
	beego.Router("/datasource/index", &controllers.DatasourceController{}, "*:Index")           //列表首页
	beego.Router("/datasource/all", &controllers.DatasourceController{}, "*:All")               //列出所有数据
	beego.Router("/datasource/updatePage", &controllers.DatasourceController{}, "*:UpdatePage") //更新页面
	beego.Router("/datasource/update", &controllers.DatasourceController{}, "*:Update")         //更新内容
	beego.Router("/datasource/addPage", &controllers.DatasourceController{}, "*:AddPage")       //新增页面
	beego.Router("/datasource/add", &controllers.DatasourceController{}, "*:Add")               //新增内容
	beego.Router("/datasource/check", &controllers.DatasourceController{}, "*:Check")           //检测数据源有效性校验
	beego.Router(models.UrlForDatasource, &controllers.DatasourceController{}, "*:GetAllDatasourceForTree")
	beego.Router(models.UrlForSchema, &controllers.DatasourceController{}, "*:GetAllSchemaForTree")
	beego.Router(models.UrlForTable, &controllers.DatasourceController{}, "*:GetAllTableForTree")
	beego.Router(models.UrlForColumn, &controllers.DatasourceController{}, "*:GetAllColumnForTree")

	//指标计算业务表
	beego.Router("/sdtBdiBusi/index", &controllers.SdtBdiBusiController{}, "*:Index")           //列表首页
	beego.Router("/sdtBdiBusi/all", &controllers.SdtBdiBusiController{}, "*:All")               //列出所有数据
	beego.Router("/sdtBdiBusi/addPage", &controllers.SdtBdiBusiController{}, "*:AddPage")       //新增页面
	beego.Router("/sdtBdiBusi/add", &controllers.SdtBdiBusiController{}, "*:Add")               //新增内容
	beego.Router("/sdtBdiBusi/updatePage", &controllers.SdtBdiBusiController{}, "*:UpdatePage") //更新页面
	beego.Router("/sdtBdiBusi/update", &controllers.SdtBdiBusiController{}, "*:Update")         //更新内容
	beego.Router("/sdtBdiBusi/delete", &controllers.SdtBdiBusiController{}, "*:Delete")         //删除内容


	beego.Router("/sdtBdiBusi/selectedTables", &controllers.SdtBdiBusiController{}, "*:SelectedTables")

	beego.Router("/sdtBdiBusi/detailConfigPage", &controllers.SdtBdiBusiController{}, "*:DetailConfigPage") //详细配置
	beego.Router("/sdtBdiBusi/processType", &controllers.SdtBdiBusiController{}, "*:ProcessType")
	beego.Router("/sdtBdiBusi/detailConfig/update", &controllers.SdtBdiBusiController{}, "*:UpdateDetailConfig") //详细配置

	// sdt_bdi_busi_config
	beego.Router("/sdtBdiBusiConfig/index", &controllers.SdtBdiBusiConfigController{}, "*:Index")                     //列表首页
	beego.Router("/sdtBdiBusiConfig/all", &controllers.SdtBdiBusiConfigController{}, "*:All")                         //列出所有数据
	beego.Router("/sdtBdiBusiConfig/updatePage", &controllers.SdtBdiBusiConfigController{}, "*:UpdatePage")           //更新页面
	beego.Router("/sdtBdiBusiConfig/update", &controllers.SdtBdiBusiConfigController{}, "*:Update")                   //更新内容
	beego.Router("/sdtBdiBusiConfig/content/synchronize", &controllers.SdtBdiBusiConfigController{}, "*:Synchronize") //同步内容
	beego.Router("/sdtBdiBusiConfig/rowMoveUp", &controllers.SdtBdiBusiConfigController{}, "*:RowMoveUp")             //行上移
	beego.Router("/sdtBdiBusiConfig/rowMoveDown", &controllers.SdtBdiBusiConfigController{}, "*:RowMoveDown")         //行下移

	beego.Router("/sdtBdiBusiFields/index", &controllers.SdtBdiBusiFieldsController{}, "*:Index")                               //列表首页
	beego.Router("/sdtBdiBusiFields/all", &controllers.SdtBdiBusiFieldsController{}, "*:All")                                   //列出所有数据
	beego.Router("/sdtBdiBusiFields/columnSelectTreePage", &controllers.SdtBdiBusiFieldsController{}, "*:ColumnSelectTreePage") //选择表字段
	beego.Router("/sdtBdiBusiFields/processFields", &controllers.SdtBdiBusiFieldsController{}, "*:ProcessFields")
	beego.Router("/sdtBdiBusiFields/processFields/add", &controllers.SdtBdiBusiFieldsController{}, "*:ProcessFieldsAdd")
	beego.Router("/sdtBdiBusiFields/processType", &controllers.SdtBdiBusiFieldsController{}, "*:ProcessType")
	beego.Router("/sdtBdiBusiFields/addFieldsPage", &controllers.SdtBdiBusiFieldsController{}, "*:AddFieldsPage")                     //新增字段首页
	beego.Router("/sdtBdiBusiFields/addConstFields/add", &controllers.SdtBdiBusiFieldsController{}, "*:AddConstFields")               //新增Const字段保存
	beego.Router("/sdtBdiBusiFields/addFields/add", &controllers.SdtBdiBusiFieldsController{}, "*:AddFields")                         //新增字段保存
	beego.Router("/sdtBdiBusiFields/sdtBdiBusiFieldsForList", &controllers.SdtBdiBusiFieldsController{}, "*:SdtBdiBusiFieldsForList") //新增字段保存
	beego.Router("/sdtBdiBusiFields/delete", &controllers.SdtBdiBusiFieldsController{}, "*:Delete")         //删除内容


	beego.Router("/sdtBdiBusiFields/rowMoveUp", &controllers.SdtBdiBusiFieldsController{}, "*:RowMoveUp")             //行上移
	beego.Router("/sdtBdiBusiFields/rowMoveDown", &controllers.SdtBdiBusiFieldsController{}, "*:RowMoveDown")         //行下移
	beego.Router("/sdtBdiBusiFields/checkData", &controllers.SdtBdiBusiFieldsController{}, "*:CheckData")             //
	beego.Router("/sdtBdiBusiFields/content/synchronize", &controllers.SdtBdiBusiFieldsController{}, "*:Synchronize") //同步内容

	//指标结果表配置
	beego.Router("/sdtBdiResult/index", &controllers.SdtBdiResultController{}, "*:Index")                   //列表首页
	beego.Router("/sdtBdiResult/all", &controllers.SdtBdiResultController{}, "*:All")                       //列出所有数据
	beego.Router("/sdtBdiResult/addPage", &controllers.SdtBdiResultController{}, "*:AddPage")               //新增页面
	beego.Router("/sdtBdiResult/add", &controllers.SdtBdiResultController{}, "*:Add")                       //新增内容
	beego.Router("/sdtBdiResult/updatePage", &controllers.SdtBdiResultController{}, "*:UpdatePage")         //更新页面
	beego.Router("/sdtBdiResult/update", &controllers.SdtBdiResultController{}, "*:Update")                 //更新内容
	beego.Router("/sdtBdiResult/datasourceTree", &controllers.SdtBdiResultController{}, "*:DatasourceTree") //更新内容

	//指标结果表字段配置
	beego.Router("/sdtBdiResultFields/index", &controllers.SdtBdiResultFieldsController{}, "*:Index")     //列表首页
	beego.Router("/sdtBdiResultFields/all", &controllers.SdtBdiResultFieldsController{}, "*:All")         //列出所有数据
	beego.Router("/sdtBdiResultFields/addPage", &controllers.SdtBdiResultFieldsController{}, "*:AddPage") //新增页面
	beego.Router("/sdtBdiResultFields/add", &controllers.SdtBdiResultController{}, "*:Add")               //新增内容
	// beego.Router("/sdtBdiResultFields/updatePage", &controllers.SdtBdiResultFieldsController{}, "*:UpdatePage") //更新页面
	// beego.Router("/sdtBdiResultFields/update", &controllers.SdtBdiResultFieldsController{}, "*:Update")         //更新内容

	//指标业务规则配置
	beego.Router("/sdtBdiBusiRule/bdi/index", &controllers.SdtBdiBusiRuleController{}, "*:BdiIndex")                        //列表首页
	beego.Router("/sdtBdiBusiRule/bdi/all", &controllers.SdtBdiController{}, "*:All")                                       //列出所有数据
	beego.Router("/sdtBdiBusiRule/columnSelectTreePage", &controllers.SdtBdiBusiRuleController{}, "*:ColumnSelectTreePage") //选择表字段
	beego.Router("/sdtBdiBusiRule/addColumn", &controllers.SdtBdiBusiRuleController{}, "*:AddColumn")                       //选择表字段
	beego.Router("/sdtBdiBusiRule/index", &controllers.SdtBdiBusiRuleController{}, "*:Index")                               //列表首页
	beego.Router("/sdtBdiBusiRule/all", &controllers.SdtBdiBusiRuleController{}, "*:All")                                   //列表首页
	beego.Router("/sdtBdiBusiRule/processFields", &controllers.SdtBdiBusiRuleController{}, "*:ProcessFields")               //列表首页
	beego.Router("/sdtBdiBusiRule/processFields/update", &controllers.SdtBdiBusiRuleController{}, "*:Update")               //列表首页

	/*
		//规则集
		beego.Router("/sdtBdiRuleSet/index", &controllers.SdtBdiRuleSetController{}, "*:Index")             //列表首页
		beego.Router("/sdtBdiRuleSet/all", &controllers.SdtBdiRuleSetController{}, "*:All")                 //列出所有数据
		beego.Router("/sdtBdiRuleSet/add", &controllers.SdtBdiRuleSetController{}, "*:Add")                 //新增内容
		beego.Router("/sdtBdiRuleSet/update", &controllers.SdtBdiRuleSetController{}, "*:Update")           //更新内容
		beego.Router("/sdtBdiRuleSet/rowMoveUp", &controllers.SdtBdiRuleSetController{}, "*:RowMoveUp")     //行上移
		beego.Router("/sdtBdiRuleSet/rowMoveDown", &controllers.SdtBdiRuleSetController{}, "*:RowMoveDown") //行下移
		beego.Router("/sdtBdiRuleSet/delete", &controllers.SdtBdiRuleSetController{}, "*:Delete")           //删除

		//规则
		beego.Router("/sdtBdiRule/index", &controllers.SdtBdiRuleController{}, "*:Index")           //列表首页
		beego.Router("/sdtBdiRule/all", &controllers.SdtBdiRuleController{}, "*:All")               //列出所有数据
		beego.Router("/sdtBdiRule/addPage", &controllers.SdtBdiRuleController{}, "*:AddPage")       //新增页面
		beego.Router("/sdtBdiRule/add", &controllers.SdtBdiRuleController{}, "*:Add")               //新增内容
		beego.Router("/sdtBdiRule/updatePage", &controllers.SdtBdiRuleController{}, "*:UpdatePage") //更新页面
		beego.Router("/sdtBdiRule/update", &controllers.SdtBdiRuleController{}, "*:Update")         //更新内容

		//规则语言
		beego.Router("/sdtBdiRuleLanguage/all", &controllers.SdtBdiRuleLanguageController{}, "*:All")         //列出所有数据
		beego.Router("/sdtBdiRuleLanguage/addPage", &controllers.SdtBdiRuleLanguageController{}, "*:AddPage") //新增页面
		beego.Router("/sdtBdiRuleLanguage/add", &controllers.SdtBdiRuleLanguageController{}, "*:Add")         //新增内容

		//KPI
		beego.Router("/sdtBdiCfgKpi/index", &controllers.SdtBdiCfgKpiController{}, "*:Index")           //列表首页
		beego.Router("/sdtBdiCfgKpi/all", &controllers.SdtBdiCfgKpiController{}, "*:All")               //列出所有数据
		beego.Router("/sdtBdiCfgKpi/addPage", &controllers.SdtBdiCfgKpiController{}, "*:AddPage")       //新增页面
		beego.Router("/sdtBdiCfgKpi/add", &controllers.SdtBdiCfgKpiController{}, "*:Add")               //新增内容
		beego.Router("/sdtBdiCfgKpi/updatePage", &controllers.SdtBdiCfgKpiController{}, "*:UpdatePage") //更新页面
		beego.Router("/sdtBdiCfgKpi/update", &controllers.SdtBdiCfgKpiController{}, "*:Update")         //更新内容
	*/

}
