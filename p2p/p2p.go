package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/justindwlee/bitcoinClone/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(w http.ResponseWriter, r *http.Request){
	// Port :3000 will upgrade the request from :4000
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	utils.HandleErr(err)
	initPeer(conn, "asdf", "Asdf")
}

func AddPeer(address string, port string){
	// Port :4000 is requesting an upgrade from the port :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws", address, port), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}