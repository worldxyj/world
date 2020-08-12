package Admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-playground/validator/v10"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"world/core"
	"world/models"
)

type MenuController struct {
	BaseController
}

func (this *MenuController) Index() {
	var menus []*models.Menu
	core.DB.Order("sort").Find(&menus)
	data := make([]interface{}, 0, 1000)
	this.format(&data, menus, "— ", 0, 0, 0)
	this.Data["Data"] = data
	this.Data["Xsrftoken"] = this.XSRFToken()
	this.TplName = "admin/menu/index.html"
}

func (this *MenuController) format(data *[]interface{}, menus []*models.Menu, leftHtml string, pid uint, level uint8, leftPadding int) {
	var temp map[string]interface{}
	for _, v := range menus {
		if v.Pid == pid {
			temp = make(map[string]interface{})
			temp["name"] = v.Name
			temp["model"] = v
			temp["level"] = level + 1
			temp["leftPadding"] = leftPadding + 0
			temp["leftHtml"] = strings.Repeat(leftHtml, int(level))
			*data = append(*data, temp)
			this.format(data, menus, leftHtml, v.Id, level+1, leftPadding+20)
		}
	}
}

func (this *MenuController) Add() {
	name := this.GetString("name")
	pid, _ := this.GetInt("pid")
	css := this.GetString("css")
	url := this.GetString("url")
	sort, _ := this.GetInt("sort")
	status, _ := this.GetInt("status")
	type AddValidate struct {
		Name   string `validate:"required,max=16" vmsg:"用户名必填且不能超过16个字符"`
		Css    string `validate:"max=16" vmsg:"css不能超过16个字符"`
		Url    string `validate:"required,max=32" vmsg:"路径必填且不能超过32个字符"`
		Pid    uint   `validate:"numeric" vmsg:"请选择正确的父级"`
		Sort   uint16 `validate:"numeric,max=60000" vmsg:"请输入正确的排序"`
		Status uint8  `validate:"oneof=0 1" vmsg:"请选择正确的状态"`
	}
	add := AddValidate{
		Name:   name,
		Css:    css,
		Url:    url,
		Pid:    uint(pid),
		Sort:   uint16(sort),
		Status: uint8(status),
	}
	err := validator.New().Struct(&add)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(&add).Elem().FieldByName(err.Field())
			vmsg := field.Tag.Get("vmsg")
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": vmsg}
			this.ServeJSON()
			this.StopRun()
		}
	}
	menu := models.Menu{
		Name:   name,
		Css:    css,
		Url:    url,
		Pid:    uint(pid),
		Sort:   uint16(sort),
		Status: uint8(status),
	}
	result := core.DB.Create(&menu)
	if result.Error == nil {
		this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "添加成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "添加失败"}
	}
	this.ServeJSON()
}

func (this *MenuController) Edit() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	menu := models.Menu{Id: uint(id)}
	result := core.DB.First(&menu).Error
	if result != nil {
		this.Data["content"] = "server error"
		this.Abort("500")
	}
	if this.Ctx.Request.Method == "POST" {
		name := this.GetString("name")
		pid, _ := this.GetInt("pid")
		css := this.GetString("css")
		url := this.GetString("url")
		sort, _ := this.GetInt("sort")
		status, _ := this.GetInt("status")
		type EditValidate struct {
			Name   string `validate:"required,max=16" vmsg:"用户名必填且不能超过16个字符"`
			Css    string `validate:"max=16" vmsg:"css不能超过16个字符"`
			Url    string `validate:"required,max=32" vmsg:"路径必填且不能超过32个字符"`
			Pid    uint   `validate:"numeric" vmsg:"请选择正确的父级"`
			Sort   uint16 `validate:"numeric,max=60000" vmsg:"请输入正确的排序"`
			Status uint8  `validate:"oneof=0 1" vmsg:"请选择正确的状态"`
		}
		edit := EditValidate{
			Name:   name,
			Css:    css,
			Url:    url,
			Pid:    uint(pid),
			Sort:   uint16(sort),
			Status: uint8(status),
		}
		err := validator.New().Struct(&edit)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				field, _ := reflect.TypeOf(&edit).Elem().FieldByName(err.Field())
				vmsg := field.Tag.Get("vmsg")
				this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": vmsg}
				this.ServeJSON()
				this.StopRun()
			}
		}
		menu := models.Menu{
			Id:     uint(id),
			Name:   name,
			Css:    css,
			Url:    url,
			Pid:    uint(pid),
			Sort:   uint16(sort),
			Status: uint8(status),
		}
		result := core.DB.Save(&menu).Error
		if result == nil {
			this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "修改成功"}
		} else {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "修改失败"}
		}
		this.ServeJSON()
	} else {
		var menus []*models.Menu
		core.DB.Order("sort").Find(&menus)
		menusTemp := make([]interface{}, 0, 1000)
		this.format(&menusTemp, menus, "— ", 0, 0, 0)
		this.Data["Menus"] = menusTemp
		this.Data["Data"] = menu
		this.Data["Xsrftoken"] = template.HTML(this.XSRFFormHTML())
		this.TplName = "admin/menu/edit.html"
	}
}

func (this *MenuController) Del() {
	id, _ := this.GetInt("id")
	menu := models.Menu{Id: uint(id)}
	result := core.DB.First(&menu).Error
	if result != nil {
		this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "未查询到记录"}
		this.ServeJSON()
		this.StopRun()
	}
	//删除子菜单
	var pMenus []*models.Menu
	var ppMenus []*models.Menu
	core.DB.Where("pid = ?", id).Find(&pMenus)
	for _, v := range pMenus {
		core.DB.Where("pid = ?", v.Id).Find(&ppMenus)
		for _, vv := range ppMenus {
			core.DB.Where("pid = ?", vv.Id).Delete(models.Menu{})
		}
		core.DB.Where("pid = ?", v.Id).Delete(models.Menu{})
	}
	core.DB.Where("pid = ?", id).Delete(models.Menu{})
	//删除菜单
	result = core.DB.Delete(&menu).Error
	if result == nil {
		this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "删除失败"}
	}
	this.ServeJSON()
}

func (this *MenuController) State() {
	id, _ := this.GetInt("id")
	menu := models.Menu{Id: uint(id)}
	core.DB.First(&menu)
	var status uint8
	var msg string
	if menu.Status == 1 {
		status = 0
		msg = "隐藏"
	} else {
		status = 1
		msg = "显示"
	}
	result := core.DB.Model(&menu).Update("status", status)
	if result.Error == nil {
		this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": msg}
	} else {
		this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "操作失败"}
	}
	this.ServeJSON()
}

func (this *MenuController) Order() {
	prefix := beego.AppConfig.String("db_dt_prefix")
	var sqlWhenThen string
	param := this.Ctx.Request.Form
	delete(param, "_xsrf")
	r := strings.NewReplacer(`'`, `\'`, `"`, `\"`, `\`, `\\`)
	for k, v := range this.Ctx.Request.Form {
		temp := " when " + r.Replace(k) + " then " + r.Replace(v[0])
		sqlWhenThen += temp
	}
	table := prefix + "menu"
	sql := fmt.Sprintf("update %s set sort = case id %s end", table, sqlWhenThen)
	core.DB.Exec(sql)
	this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "操作成功"}
	this.ServeJSON()
}
