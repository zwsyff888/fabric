package controllers

import (
	// "crypto/tls"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/config"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	// "fmt"
	"github.com/hyperledger/fabric/bddtests/fabricBrowser/fabricBrowser/models"
	// "strconv"
	// "yxchainExplorer_0.6.2/models"
	// "math/rand"
	// "time"
)

type NewController struct {
	beego.Controller
}

func (this *NewController) GetChainID() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")

	if !models.BaseCheck() {
		this.Data["json"] = []string{}
	} else {
		this.Data["json"] = models.GetChainIDs()
	}
	this.ServeJSON()

}

func (this *NewController) GetPeers() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	chainid := this.GetString("chainid")
	if !models.BaseCheck(chainid) {
		this.Data["json"] = []string{}
	} else {

		this.Data["json"] = models.GetPeerStatusMap(chainid)

	}
	this.ServeJSON()
}

func (this *NewController) GetBlocks() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")

	tmpNum, err := this.GetInt("blocknum")
	chainid := this.GetString("chainid")

	if err != nil || tmpNum < 0 {
		tmpNum = 0
	}
	var startNum uint64
	startNum = uint64(tmpNum)

	if !models.BaseCheck(chainid) {
		this.Data["json"] = []string{}
	} else {
		this.Data["json"] = models.GetBlocks(startNum, chainid)
	}
	this.ServeJSON()
}

/*

func (this *NewController) GetRate() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")

	if !models.BaseCheck() {
		this.Data["json"] = []string{}
	} else {
		this.Data["json"] = models.GetBlocks(startNum)
	}

	this.Data["json"] = models.GetRate()
	this.ServeJSON()
}*/

func (this *NewController) GetBlockInfo() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	// r := this.Ctx.Input.Param(":splat")
	// fmt.Println("chenqiao r: ", r)
	chainid := this.GetString("chainid")
	tmpNum, err := this.GetInt("startnum")
	if err != nil || tmpNum < 0 {
		tmpNum = 0
	}

	if !models.BaseCheck(chainid) {
		this.Data["json"] = []string{}
	} else {

		peerid, _ := models.GetMaxPeer(chainid)

		blockNum := uint64(tmpNum)

		this.Data["json"] = models.GetBlockinfos(peerid, blockNum, chainid)

	}
	this.ServeJSON()

}

func (this *NewController) GetBlockDetail() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	// peerIP := this.GetString("peerip")
	// tmpNum, err := this.GetInt("blocknum")
	tmpNum, err := this.GetInt("blocknum")
	chainid := this.GetString("chainid")
	if err != nil || tmpNum < 0 {
		tmpNum = 0
	}

	if !models.BaseCheck(chainid) {
		this.Data["json"] = []string{}
	} else {
		peerid, _ := models.GetMaxPeer(chainid)

		blockNum := uint64(tmpNum)
		// fmt.Println("blocknum ", blockNum)
		tmp := models.GetBlockDetail(peerid, blockNum, chainid)
		if tmp != nil {
			this.Data["json"] = *tmp
		} else {
			this.Data["json"] = []string{}
		}

	}
	this.ServeJSON()
	// r := this.Ctx.Input.Param(":splat")
	// _, err := strconv.Atoi(r)
	// if err != nil {
	// 	this.Data["json"] = []string{}
	// 	this.ServeJSON()
	// 	return
	// }
	// peers := models.BlocksHandler.GetPeers()
	// maxHeight := 0
	// addr := peers[0].Address
	// for _, p := range peers {
	// 	if p.Height > maxHeight {
	// 		maxHeight = p.Height
	// 		addr = p.Address
	// 	}
	// }

	// prefix := "http"
	// peerCfg, err := config.NewConfig("json", "conf/config.json")
	// if err == nil {
	// 	tlsEnable, _ := peerCfg.Bool("tlsEnabled")
	// 	if tlsEnable {
	// 		prefix = "https"
	// 	}
	// }

	// httpsUrl := prefix + "://" + addr + "/chain/blocks/" + r

	// req := httplib.Get(httpsUrl).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	// str, _ := req.String()
	// this.Ctx.WriteString(str)
}
