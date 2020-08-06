package Home

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	c.TplName = "home/index/index.html"
}
