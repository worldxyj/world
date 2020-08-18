package Admin

import (
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"strings"
	"world/core"
	"world/models"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) Index() {
	var (
		roles    []*models.Role
		listRows = 15
		p        = 1
		count    int
		err      error
	)
	params := core.Path2Map(this.Ctx.Input.Param(":splat"))
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
	core.DB.Find(&roles).Count(&count)
	core.DB.Offset(listRows * (p - 1)).Limit(listRows).Order("id desc").Find(&roles)
	this.Data["Page"] = core.Paginations(6, count, listRows, p, beego.URLFor("RoleController.Index"))
	this.Data["Data"] = roles
	this.Data["Xsrftoken"] = this.XSRFToken()
	this.TplName = "admin/role/index.html"
}

func (this *RoleController) Add() {
	if this.Ctx.Request.Method == "POST" {
		name := this.GetString("name")
		menuIds := this.GetStrings("menu_ids[]")
		if name == "" || len(name) > 32 {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "角色名称不能为空且不能超过32个字符"}
			this.ServeJSON()
			this.StopRun()
		}
		if len(menuIds) == 0 {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "请选择菜单"}
			this.ServeJSON()
			this.StopRun()
		}
		role := models.Role{Name: name}
		tx := core.DB.Begin()
		result := tx.Create(&role).Error
		if result != nil {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "添加失败"}
			this.ServeJSON()
			this.StopRun()
		}
		var roleMenu models.RoleMenu
		var menuId int
		for _, v := range menuIds {
			menuId, _ = strconv.Atoi(v)
			roleMenu = models.RoleMenu{
				RoleId: role.Id,
				MenuId: uint(menuId),
			}
			result = tx.Create(&roleMenu).Error
			if result != nil {
				tx.Rollback()
				this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "添加失败"}
				this.ServeJSON()
				this.StopRun()
			}
		}
		tx.Commit()
		this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "添加成功"}
		this.ServeJSON()
	} else {
		var menus []*models.Menu
		core.DB.Order("sort").Find(&menus)
		var data []interface{}
		this.format(&data, menus, "— ", 0, 0, 0)
		this.Data["Data"] = data
		this.Data["Xsrftoken"] = template.HTML(this.XSRFFormHTML())
		this.TplName = "admin/role/add.html"
	}
}

func (this *RoleController) format(data *[]interface{}, menus []*models.Menu, leftHtml string, pid uint, level uint8, leftPadding int) {
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

func (this *RoleController) Edit() {
	id, _ := this.GetInt(":id")
	role := models.Role{}
	result := core.DB.Where("id = ?", id).First(&role).Error
	if result != nil {
		this.Abort("404")
	}
	if this.Ctx.Request.Method == "POST" {
		name := this.GetString("name")
		menuIds := this.GetStrings("menu_ids[]")
		if name == "" || len(name) > 32 {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "角色名称不能为空且不能超过32个字符"}
			this.ServeJSON()
			this.StopRun()
		}
		if len(menuIds) == 0 {
			this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "请选择菜单"}
			this.ServeJSON()
			this.StopRun()
		}
		core.DB.Model(&role).Update("name", name)
		core.DB.Where("role_id = ?", id).Delete(&models.RoleMenu{})
		var roleMenu models.RoleMenu
		var menuId int
		for _, v := range menuIds {
			menuId, _ = strconv.Atoi(v)
			roleMenu = models.RoleMenu{
				RoleId: role.Id,
				MenuId: uint(menuId),
			}
			core.DB.Create(&roleMenu)
		}
		this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "修改成功"}
		this.ServeJSON()
	} else {
		var menuId []uint
		core.DB.Model(&models.RoleMenu{}).Where("role_id = ?", id).Pluck("menu_id", &menuId)
		var menus []*models.Menu
		core.DB.Order("sort").Find(&menus)
		var data []interface{}
		this.format(&data, menus, "— ", 0, 0, 0)
		this.Data["role"] = role
		this.Data["menuId"] = menuId
		this.Data["data"] = data
		this.Data["xsrftoken"] = template.HTML(this.XSRFFormHTML())
		this.TplName = "admin/role/edit.html"
	}
}

func (this *RoleController) Del() {
	id, _ := this.GetInt("id")
	result := core.DB.Where("id = ?", id).Delete(models.Role{}).Error
	if result != nil {
		this.Data["json"] = map[string]interface{}{"errcode": 1, "msg": "删除失败"}
		this.ServeJSON()
		this.StopRun()
	}
	core.DB.Where("role_id = ?", id).Delete(models.RoleMenu{})
	this.Data["json"] = map[string]interface{}{"errcode": 0, "msg": "删除成功"}
	this.ServeJSON()
}
