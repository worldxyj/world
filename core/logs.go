package core

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

func logsInit() {
	logConf := make(map[string]interface{})
	logConf["filename"] = beego.AppConfig.String("logfilename")
	logConf["separate"] = strings.Split(beego.AppConfig.String("logseparate"), ",")
	if beego.AppConfig.String("runmode") == "dev" {
		logConf["level"] = 1
	} else {
		logConf["level"], _ = beego.AppConfig.Int("loglevel")
	}
	logConf["daily"] = true
	logConf["maxdays"] = 7
	confJson, _ := json.Marshal(logConf)
	logs.SetLogger(logs.AdapterMultiFile, string(confJson))
	logs.EnableFuncCallDepth(true)
	logs.Async()
}
