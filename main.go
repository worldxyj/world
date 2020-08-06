package main

import (
	"github.com/astaxie/beego"
	_ "world/core"
	_ "world/routers"
)

func main() {
	beego.Run()
}
