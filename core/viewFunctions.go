package core

import (
	"github.com/astaxie/beego"
	"strings"
	"time"
	"world/models"
)

func viewFunctions() {
	beego.AddFuncMap("date", date)
	beego.AddFuncMap("menu", menu)
}

func date(intTime uint) string {
	return time.Unix(int64(intTime), 0).Format("2006-01-02 15:04:05")
}

func menu(uri string) map[string]interface{} {
	var menus []*models.Menu
	DB.Where("status = 1").Order("sort").Find(&menus)
	uriSlice := strings.Split(uri, "/")
	var first, second, third string
	if len(uriSlice) > 2 && len(uriSlice) < 5 {
		first = uriSlice[2]
		second = first + "/" + uriSlice[3]
	} else if len(uriSlice) > 4 {
		first = uriSlice[2]
		second = first + "/" + uriSlice[3]
		third = second + "/" + uriSlice[4]
	} else {
		first = ""
		second = ""
		third = ""
	}
	//三级菜单
	var dataNav []*models.Menu
	for _, v := range menus {
		if v.Url == second {
			for _, vv := range menus {
				if vv.Pid == v.Id {
					dataNav = append(dataNav, vv)
				}
			}
		}
	}
	return map[string]interface{}{
		"data":    menus,
		"dataNav": dataNav,
		"first":   first,
		"second":  second,
		"third":   third,
	}
}
