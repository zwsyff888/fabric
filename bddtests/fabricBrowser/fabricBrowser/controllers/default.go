package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/hyperledger/fabric/bddtests/fabricBrowser/fabricBrowser/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	// c.Data["json"] = models.AllChannelPeerStatusMap
	// c.ServeJSON()
}

type SuperviseController struct {
	beego.Controller
}

func (c *SuperviseController) Get() {
	fmt.Println("lslsllsls")
	c.TplName = "supervise.html"
}

type ZhController struct {
	beego.Controller
}

func (c *ZhController) Get() {
	c.TplName = "index-zh/index.html"
}
