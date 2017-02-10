package routers

import (
	"github.com/hyperledger/fabric/bddtests/supervise/server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
