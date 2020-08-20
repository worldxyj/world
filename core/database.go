package core

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

var DB *gorm.DB

func databaseInit() {
	dbType := beego.AppConfig.String("db_type")
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	dbName := beego.AppConfig.String(dbType + "::db_name")
	dbCharset := beego.AppConfig.String(dbType + "::db_charset")
	var openConf string
	if dbType == "mysql" {
		openConf = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName, dbCharset)
	} else if dbType == "postgres" {
		openConf = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPwd)
	} else if dbType == "sqlite3" {
		openConf = dbName
	} else {
		panic("数据库类型错误")
	}

	var err error
	DB, err = gorm.Open(dbType, openConf)
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
	//UpdateAt修改为unix时间戳
	DB.Callback().Update().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if _, ok := scope.Get("gorm:update_column"); !ok {
			_ = scope.SetColumn("UpdatedAt", time.Now().Unix())
		}
	})
	//CreatedAt修改为unix时间戳
	DB.Callback().Create().Replace("gorm:update_time_stamp", func(scope *gorm.Scope) {
		if !scope.HasError() {
			if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
				if createdAtField.IsBlank {
					_ = createdAtField.Set(time.Now().Unix())
				}
			}
			if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
				if updatedAtField.IsBlank {
					_ = updatedAtField.Set(time.Now().Unix())
				}
			}
		}
	})
	//开启日志
	dbLog, _ := beego.AppConfig.Bool("db_log")
	if dbLog {
		DB.LogMode(true)
	}
}
