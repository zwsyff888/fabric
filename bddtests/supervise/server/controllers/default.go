package controllers

import (
	"github.com/astaxie/beego"
)

/*type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}*/

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.tpl"
}
