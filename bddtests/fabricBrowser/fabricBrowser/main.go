package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/hyperledger/fabric/bddtests/fabricBrowser/fabricBrowser/models"
	_ "github.com/hyperledger/fabric/bddtests/fabricBrowser/fabricBrowser/routers"
)

func main() {
	logs.Debug("hehehehehehe")
	go models.StartServer()
	go models.CheckEnable()
	// logs.Debug("hehehehehehe")
	// go models.StartServer()
	models.PeerStatusMap = make(map[string]*models.PeerMessage)
	// models.QueryClient()

	beego.SetStaticPath("/super", "static/supervise")
	beego.Run()

}
