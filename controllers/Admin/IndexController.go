package Admin

import (
	"crypto/md5"
	"fmt"
	"github.com/go-playground/validator/v10"
	"html/template"
	"reflect"
	"world/core"
	"world/models"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	this.TplName = "admin/index/index.html"
}

func (this *IndexController) EditPwd() {
	if this.Ctx.Request.Method == "POST" {
		//表单验证
		type EditPwdValidate struct {
			Password    string `validate:"required" vmsg:"请输入原密码"`
			NewPassword string `validate:"required" vmsg:"请输入新密码"`
		}
		password := this.GetString("password")
		newPassword := this.GetString("new_password")
		if password == newPassword {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "新密码不能与原密码一样"}
			this.ServeJSON()
			this.StopRun()
		}
		editPwd := &EditPwdValidate{
			Password:    password,
			NewPassword: newPassword,
		}
		validate := validator.New()
		err := validate.Struct(editPwd)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				field, _ := reflect.TypeOf(editPwd).Elem().FieldByName(err.Field())
				vmsg := field.Tag.Get("vmsg")
				this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": vmsg}
				this.ServeJSON()
				this.StopRun()
			}
		}
		//验证原密码
		admin := models.Admin{
			Id: uint(this.AdminId),
		}
		result := core.DB.First(&admin)
		if result.Error != nil {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "没有查询到管理员"}
			this.ServeJSON()
			this.StopRun()
		}
		password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
		newPassword = fmt.Sprintf("%x", md5.Sum([]byte(newPassword)))
		if admin.Password != password {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "原密码错误"}
			this.ServeJSON()
			this.StopRun()
		}
		//修改密码
		result = core.DB.Model(&admin).Update("password", newPassword)
		if result.Error != nil {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "修改失败"}
		} else {
			this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "修改成功"}
		}
		this.ServeJSON()
		this.StopRun()
	} else {
		this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
		this.TplName = "admin/index/editPwd.html"
	}
}
