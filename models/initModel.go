package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Local"
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(User))
	// orm.RegisterModel(new(SdtBdiBase))
	orm.RegisterModel(new(AdtBdiAdm))
	//orm.RegisterModel(new(SdtBdiDomain))

	//不注册此model，使用原生sql操作增删改查
	//orm.RegisterModel(new(SdtBdiSet))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
