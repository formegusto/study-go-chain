package p2p

import (
	"fmt"
	"net/http"
	"time"

	"github.com/formegusto/study-go-chain/utils"
	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// port :3000 will upgrade the request from :4000
	address := utils.Splitter(r.RemoteAddr, ":", 0)
	openPort := r.URL.Query().Get("openPort")
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && address != ""
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)

	initPeer(conn, address, openPort)

	time.Sleep(10 * time.Second)
	conn.WriteMessage(websocket.TextMessage, []byte("Hello from Port 3000!"))
}

func AddPeer(address, port, openPort string) {
	// port :4000 is requesting an upgrade from the port :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)

	time.Sleep(10 * time.Second)
	conn.WriteMessage(websocket.TextMessage, []byte("Hello from Port 4000!"))
}