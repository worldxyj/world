package Admin

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	AdminId int
}

func (this *BaseController) Prepare() {
	var ok bool

	//模拟登陆
	this.SetSession("AdminId", 1)
	type AdminInfo1 struct {
		Id        int
		Name      string
		Password  string
		Tel       string
		Ip        string
		LoginAt   int64
		CreatedAt int64
	}
	adminInfo1 := &AdminInfo1{
		Id:       1,
		Name:     "模拟登陆",
		Password: "e10adc3949ba59abbe56e057f20f883e",
		Tel:      "13839958207",
		Ip:       "127.0.0.1",
	}
	this.SetSession("AdminInfo", adminInfo1)

	AdminId := this.GetSession("AdminId")
	this.AdminId, ok = AdminId.(int)
	if !ok || AdminId == 0 {
		this.Redirect(beego.URLFor("LoginController.Index"), 302)
		return
	}
	AdminInfo := this.GetSession("AdminInfo")
	this.Data["AdminInfo"] = AdminInfo
	//菜单状态
	this.Data["Uri"] = this.Ctx.Request.RequestURI
}
