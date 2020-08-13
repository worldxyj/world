package Admin

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-playground/validator/v10"
	"html/template"
	"reflect"
	"strconv"
	"time"
	"world/core"
	"world/models"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) Index() {
	var (
		admins   []*models.Admin
		listRows = 15
		p        = 1
		err      error
	)
	params := core.Path2Map(this.Ctx.Input.Param(":splat"))
	//页码处理
	if _, ok := params["p"]; ok {
		p, err = strconv.Atoi(params["p"])
		if err != nil {
			p = 1
		}
	} else {
		p, err = this.GetInt("p")
		if err != nil {
			p = 1
		}
	}
	keyword := this.GetString("keyword")
	var count int
	core.DB.Where(func(keyword string) (string, string) {
		if keyword != "" {
			return "name like ?", "%" + keyword + "%"
		}
		return "1 = ?", "1"
	}(keyword)).Find(&admins).Count(&count)
	core.DB.Where(func(keyword string) (string, string) {
		if keyword != "" {
			return "name like ?", "%" + keyword + "%"
		}
		return "1 = ?", "1"
	}(keyword)).Offset(listRows * (p - 1)).Limit(listRows).Order("id desc").Find(&admins)
	this.Data["Page"] = core.Paginations(6, count, listRows, p, beego.URLFor("AdminController.Index"), "keyword", params["keyword"])
	this.Data["Data"] = &admins
	this.Data["Xsrftoken"] = this.XSRFToken()
	this.Data["Keyword"] = keyword
	this.TplName = "admin/admin/index.html"
}

func (this *AdminController) Add() {
	if this.Ctx.Request.Method == "POST" {
		name := this.GetString("name")
		password := this.GetString("password")
		tel := this.GetString("tel")
		type AddValidate struct {
			Name     string `validate:"required,max=20" vmsg:"请输入正确的用户名"`
			Password string `validate:"required,min=6" vmsg:"密码不能小于6位"`
			Tel      string `validate:"max=11" vmsg:"手机号不能超过11位"`
		}
		add := &AddValidate{
			Name:     name,
			Password: password,
			Tel:      tel,
		}
		err := validator.New().Struct(add)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				field, _ := reflect.TypeOf(add).Elem().FieldByName(err.Field())
				vmsg := field.Tag.Get("vmsg")
				this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": vmsg}
				this.ServeJSON()
				this.StopRun()
			}
		}
		admin := models.Admin{
			Name:      name,
			Password:  fmt.Sprintf("%x", md5.Sum([]byte(password))),
			Tel:       tel,
			CreatedAt: uint(time.Now().Unix()),
		}
		result := core.DB.Create(&admin)
		if result.Error == nil {
			this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "添加成功"}
		} else {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "添加失败"}
		}
		this.ServeJSON()
		this.StopRun()
	} else {
		this.Data["Xsrfdata"] = template.HTML(this.XSRFFormHTML())
		this.TplName = "admin/admin/add.html"
	}
}

func (this *AdminController) Edit() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if this.Ctx.Request.Method == "POST" {
		name := this.GetString("name")
		password := this.GetString("password")
		tel := this.GetString("tel")
		type EditValidate struct {
			Name     string `validate:"required,max=20" vmsg:"请输入正确的用户名"`
			Password string `validate:"omitempty,min=6" vmsg:"密码不能小于6位"`
			Tel      string `validate:"max=11" vmsg:"请输入正确的手机号"`
		}
		edit := &EditValidate{
			Name:     name,
			Password: password,
			Tel:      tel,
		}
		err := validator.New().Struct(edit)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				field, _ := reflect.TypeOf(edit).Elem().FieldByName(err.Field())
				vmsg := field.Tag.Get("vmsg")
				this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": vmsg}
				this.ServeJSON()
				this.StopRun()
			}
		}
		admin := models.Admin{
			Id: uint(id),
		}
		result := core.DB.First(&admin)
		if result.Error != nil {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "未查询到记录"}
			this.ServeJSON()
			this.StopRun()
		}
		data := make(map[string]interface{})
		data["name"] = name
		data["tel"] = tel
		if password != "" {
			data["password"] = fmt.Sprintf("%x", md5.Sum([]byte(password)))
		}
		result = core.DB.Model(&admin).Updates(data)
		if result.Error == nil {
			this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "修改成功"}
		} else {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "修改失败"}
		}
		this.ServeJSON()
	} else {
		admin := models.Admin{
			Id: uint(id),
		}
		result := core.DB.First(&admin)
		if result.Error != nil {
			this.Data["content"] = "server error"
			this.Abort("500")
		}
		this.Data["Data"] = &admin
		this.Data["Xsrfdata"] = template.HTML(this.XSRFFormHTML())
		this.TplName = "admin/admin/edit.html"
	}
}

func (this *AdminController) Del() {
	id, _ := this.GetInt("id")
	if id == 1 {
		this.Data["json"] = map[string]interface{}{
			"errcode": 1,
			"msg":     "不能删除admin",
		}
		this.ServeJSON()
		this.StopRun()
	}
	admin := models.Admin{
		Id: uint(id),
	}
	result := core.DB.Where("id = ?", id).First(&admin)
	if result.Error != nil {
		this.Data["json"] = map[string]interface{}{
			"errcode": 1,
			"msg":     "未查询到记录",
		}
		this.ServeJSON()
		this.StopRun()
	}
	result = core.DB.Delete(&admin)
	if result.Error != nil {
		this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "删除失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "删除成功"}
	}
	this.ServeJSON()
	this.StopRun()
}
