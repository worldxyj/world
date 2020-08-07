package Admin

import (
	"fmt"
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
	this.format(&data, menus, "- ", 0, 0, 0)
	fmt.Println(data)
}

func (this *MenuController) format(data *[]interface{}, menus []*models.Menu, leftHtml string, pid uint, level uint8, leftPadding int) {
	var temp map[string]interface{}
	temp = make(map[string]interface{})
	for _, v := range menus {
		if v.Pid == pid {
			//temp = make(map[string]interface{})
			temp["name"] = v.Name
			temp["model"] = v
			temp["level"] = level + 1
			temp["leftPadding"] = leftPadding + 0
			temp["leftHtml"] = strings.Repeat(leftHtml, int(level))
			*data = append(*data, temp)
			fmt.Println(temp)
			fmt.Println(data)
			this.format(data, menus, leftHtml, v.Id, level+1, leftPadding+20)
		}
	}
}
