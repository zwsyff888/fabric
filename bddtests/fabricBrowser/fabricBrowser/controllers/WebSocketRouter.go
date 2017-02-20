package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type WSController struct {
	beego.Controller
	// i18n.Locale
}

func (this *WSController) Test() {
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	_, p, err := ws.ReadMessage()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("chenqiao!!!!!!:  ", p)

	ws.WriteJSON("hahahaha")
	// ws.

	this.Data["json"] = []string{}
	this.ServeJSON()
}
