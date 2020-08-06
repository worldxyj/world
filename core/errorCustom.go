package core

import (
	"github.com/astaxie/beego"
	"world/controllers"
)

func errorCustom() {
	beego.ErrorController(&controllers.ErrorController{})
}
