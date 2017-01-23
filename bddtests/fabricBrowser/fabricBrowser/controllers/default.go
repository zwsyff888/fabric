package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/hyperledger/fabric/bddtests/fabricBrowser/fabricBrowser/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	// c.Data["json"] = models.PeerStatusMap
	// c.ServeJSON()
}
