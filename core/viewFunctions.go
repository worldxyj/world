package core

import (
	"github.com/astaxie/beego"
	"time"
	"world/models"
)

func viewFunctions() {
	beego.AddFuncMap("date", date)
	beego.AddFuncMap("menu", menu)
}

func date(intTime int64) string {
	return time.Unix(intTime, 0).Format("2006-01-02 15:04:05")
}

func menu() []*models.Menu {
	var menus []*models.Menu
	DB.Where("status = 1").Order("sort").Find(&menus)
	return menus
}
