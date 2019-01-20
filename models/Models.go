package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego"
)

func RegisterDB() {
	beego.Debug("【注册数据库】")

    mysqlurls := beego.AppConfig.String("mysqlurls")
    mysqldb := beego.AppConfig.String("mysqldb")
    mysqluser := beego.AppConfig.String("mysqluser")
    mysqlpass := beego.AppConfig.String("mysqlpass")
    mysqlport := beego.AppConfig.String("mysqlport")
    
	// set default database
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":"+mysqlport+")/"+mysqldb+"?charset=utf8", 30)

	// create table
	orm.RunSyncdb("default", false, true)

}

