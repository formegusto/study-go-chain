package p2p

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/formegusto/study-go-chain/utils"
	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// port :3000 will upgrade the request from :4000

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)

	address := strings.Split(r.RemoteAddr, ":")[0]
	openPort := r.URL.Query().Get("openPort")
	initPeer(conn, address, openPort)
}

func AddPeer(address, port, openPort string) {
	// port :4000 is requesting an upgrade from the port :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}