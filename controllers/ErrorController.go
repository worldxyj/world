package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "页面不存在"
	c.TplName = "errors/404.html"
}

func (c *ErrorController) Error500() {
	//c.Data["content"] = "server error"
	c.TplName = "errors/500.html"
}
