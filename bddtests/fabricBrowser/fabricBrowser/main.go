package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	// "github.com/beego/i18n"
	"github.com/hyperledger/fabric/bddtests/fabricBrowser/fabricBrowser/models"
	_ "github.com/hyperledger/fabric/bddtests/fabricBrowser/fabricBrowser/routers"
)

func main() {
	logs.Debug("hehehehehehe!!!!")
	go models.StartServer()
	go models.CheckEnable()
	// logs.Debug("hehehehehehe")
	// go models.StartServer()
	models.AllChannelPeerStatusMap = make(map[string](map[string]*models.PeerMessage))
	// models.QueryClient()
	logs.Debug(models.AllChannelPeerStatusMap)
	beego.SetStaticPath("/super", "static/supervise")
	models.SocketsProperty = models.NewProperty()
	// beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()

}
