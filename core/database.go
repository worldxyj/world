package core

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func databaseInit() {
	dbType := beego.AppConfig.String("db_type")
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	dbName := beego.AppConfig.String(dbType + "::db_name")
	openConf := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)
	var err error
	DB, err = gorm.Open("mysql", openConf)
	if err != nil {
		panic(err)
	}
	// 不使用复数表名
	DB.SingularTable(true)
	// 设置表前缀
	dbDtPrefix := beego.AppConfig.String("db_dt_prefix")
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbDtPrefix + defaultTableName
	}
	//开启日志
	//DB.LogMode(true)
}
