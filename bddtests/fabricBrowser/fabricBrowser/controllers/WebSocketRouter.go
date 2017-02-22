package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/hyperledger/fabric/bddtests/fabricBrowser/fabricBrowser/models"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type WSController struct {
	beego.Controller
	// i18n.Locale
}

type DataJson struct {
	Chainid  string `json:"chainid"`
	BlockNum string `json:"blocknum"`
	TxID     string `json:"txid"`
}

func ReadData(ws *websocket.Conn, chain *string, blocknum *string, txid *string, ReadProperty models.Property) {
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			ws.Close()
			return
		} else {

			fmt.Println("LLLLLLLLL", string(data))
			var config DataJson
			if jsonerr := json.Unmarshal([]byte(string(data)), &config); jsonerr == nil {
				if len(config.Chainid) != 0 {
					*chain = config.Chainid
				}
				*blocknum = config.BlockNum
				*txid = config.TxID
				ReadProperty.Update()
			} else {
				ws.WriteJSON("wrong input!!!!")
			}

		}
	}

}

func SendData(ws *websocket.Conn, canSend bool, chainid string, beforeChainIDs *models.APISChainIDs, beforeBlockinfos *models.SBlockinfo, beforeBlocks *models.APISBlock, beforePeers *models.APISPeerList, beforeTransInfo *models.APISTransInfo) {
	canSend, beforeChainIDs.Data = models.APISCheckChainIDs(beforeChainIDs.Data)
	if canSend {
		ws.WriteJSON(beforeChainIDs)
	}

	canSend, beforeBlockinfos.Data = models.APISCheckBlockinfos(beforeBlockinfos.Data, chainid)

	if canSend {
		ws.WriteJSON(beforeBlockinfos)
	}

	canSend, beforeBlocks.Data = models.APISCheckSBlock(beforeBlocks.Data, chainid)

	if canSend {
		ws.WriteJSON(beforeBlocks)
	}

	canSend, beforePeers.Data = models.APISCheckSPeerList(beforePeers.Data, chainid)
	if canSend {
		ws.WriteJSON(beforePeers)
	}

	canSend, beforeTransInfo.Data = models.APISCheckTransInfo(beforeTransInfo.Data, chainid)

	if canSend {
		ws.WriteJSON(beforeTransInfo)
	}

}

func (this *WSController) Test() {
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	defer ws.Close()

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	writestream := models.SocketsProperty.Observe()
	var ReadProperty models.Property

	ReadProperty = models.NewProperty()

	readstream := ReadProperty.Observe()

	// 存储之前数据
	// 此处需要考虑锁的问题
	var DataLock sync.Mutex
	beforeChainIDs := models.NilAPISChainIDs()
	beforeBlockinfos := models.NilSBlockinfos()
	beforeBlockdetail := models.NilSMblock()
	beforeBlocks := models.NilSBlock()
	beforePeers := models.NilSPeerStatus()
	beforeTransInfo := models.NilSAPITransInfo()
	beforeTransDetail := models.NilSTransData()

	var canSend bool = false
	var chainid string = ""
	var blocknum string = ""
	var txid string = ""
	// 第一次连接 不停获取chainsID／直到取到为止
	timeout := make(chan bool, 1)

	go func() {

		time.Sleep(time.Second * 10)
		timeout <- true
	}()

	for {
		canSend, beforeChainIDs.Data = models.APISCheckChainIDs(beforeChainIDs.Data)
		if canSend {
			ws.WriteJSON(beforeChainIDs)
			fmt.Println("FFFirst input is end!!!")
			break
		}
		select {
		case <-timeout:
			ws.WriteJSON("time out, socket is close now")
			ws.Close()
			return
		}
	}

	// 等待前端的id返回

	go ReadData(ws, &chainid, &blocknum, &txid, ReadProperty)

	// 当有数据之后的处理
	for {
		select {
		//用于监听fabric端数据变化
		case <-writestream.Changes():
			{
				writestream.Next()
				DataLock.Lock()
				if len(chainid) == 0 {
					DataLock.Unlock()
					continue
				}
				//无参数的输出
				SendData(ws, canSend, chainid, beforeChainIDs, beforeBlockinfos, beforeBlocks, beforePeers, beforeTransInfo)
				DataLock.Unlock()
				fmt.Println("DATA is send OK")
			}

		//用于监听socket端数据变化
		case <-readstream.Changes():
			{
				readstream.Next()
				// 取到之后获取所有数据的初始值
				fmt.Println("client is Input!!! ", chainid)
				DataLock.Lock()
				//无参数的输出
				if chainid != "" {
					SendData(ws, canSend, chainid, beforeChainIDs, beforeBlockinfos, beforeBlocks, beforePeers, beforeTransInfo)
				}

				//有参数的输出 blockdetail
				if blocknum != "" {
					Num, numerr := strconv.Atoi(blocknum)
					if numerr != nil {
						ws.WriteJSON("wrong BlockNum!")
						return
					}
					if Num < 0 {
						Num = 0
					}
					newNum := uint64(Num)
					fmt.Println("BBBBlockNum:  ", newNum)
					fmt.Println("CCCChainid:   ", chainid)
					_, beforeBlockdetail.Data = models.APISCheckSMblock(beforeBlockdetail.Data, chainid, newNum)
					ws.WriteJSON(beforeBlockdetail)

				}

				// txdetail
				if txid != "" {
					fmt.Println("TTTTTTXID: ", txid)
					_, beforeTransDetail.Data = models.APISCheckTransData(beforeTransDetail.Data, chainid, txid)
					ws.WriteJSON(beforeTransDetail)
				}

				DataLock.Unlock()
			}

		}

		//链状态

		//块信息

	}

	// ws.

	this.Data["json"] = []string{}
	this.ServeJSON()
}
