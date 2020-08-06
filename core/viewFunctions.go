package core

import (
	"github.com/astaxie/beego"
	"time"
)

func viewFunctions() {
	beego.AddFuncMap("date", date)
}

func date(intTime int64) string {
	return time.Unix(intTime, 0).Format("2006-01-02 15:04:05")
}
