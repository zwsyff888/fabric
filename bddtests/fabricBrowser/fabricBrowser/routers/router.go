package routers

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric/bddtests/fabricBrowser/fabricBrowser/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/supervise", &controllers.SuperviseController{})
	// beego.Router("/peers", &controllers.MainController{}, "*:GetPeers")
	// beego.Router("/gethttps/*", &controllers.MainController{}, "*:GetHttps")

	//new router for blocksinfos restful api
	beego.Router("/withblocksinfos/chainids", &controllers.NewController{}, "*:GetChainID")
	beego.Router("/withblocksinfos/peers", &controllers.NewController{}, "*:GetPeers")
	beego.Router("/withblocksinfos/blocks/*", &controllers.NewController{}, "*:GetBlocks")
	// beego.Router("/withblocksinfos/rate", &controllers.NewController{}, "*:GetRate")
	beego.Router("/withblocksinfos/blockinfo/*", &controllers.NewController{}, "*:GetBlockInfo")
	beego.Router("/withblocksinfos/blockdetail/*", &controllers.NewController{}, "*:GetBlockDetail")

	beego.Router("/withblocksinfos/transinfo/*", &controllers.NewController{}, "*:GetTransInfo")
	beego.Router("/withblocksinfos/transdetail/*", &controllers.NewController{}, "*:GetTransDetail")

	beego.Router("/index-zh", &controllers.ZhController{})
	beego.Router("/ws/test", &controllers.WSController{}, "*:Test")

}
