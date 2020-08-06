package Admin

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-playground/validator/v10"
	"html/template"
	"reflect"
	"time"
	"world/core"
	"world/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Index() {
	if this.Ctx.Request.Method == "POST" {
		name := this.GetString("name")
		pwd := this.GetString("password")
		type LoginValidate struct {
			Name     string `validate:"required,max=20" vmsg:"请输入正确的用户名"`
			Password string `validate:"required" vmsg:"请输入密码"`
		}
		login := &LoginValidate{
			Name:     name,
			Password: pwd,
		}
		err := validator.New().Struct(login)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				field, _ := reflect.TypeOf(login).Elem().FieldByName(err.Field())
				vmsg := field.Tag.Get("vmsg")
				this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": vmsg}
				this.ServeJSON()
				this.StopRun()
			}
		}
		pwd = fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
		admin := models.Admin{Name: name, Password: pwd}
		if err := core.DB.Find(&admin, admin).Error; err != nil {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "用户名或密码错误"}
			this.ServeJSON()
			this.StopRun()
		}
		//记录ip和登录时间
		admin.Ip = this.Ctx.Input.IP()
		admin.LoginAt = time.Now().Unix()
		core.DB.Save(&admin)
		this.SetSession("AdminId", admin.Id)
		this.SetSession("AdminInfo", admin)
		this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "登录成功"}
		this.ServeJSON()
	} else {
		this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
		this.TplName = "admin/login/index.html"
	}
}

func (this *LoginController) Logout() {
	this.DelSession("AdminId")
	this.DelSession("AdminInfo")
	this.Redirect(beego.URLFor("AdminController.Index"), 302)
}
