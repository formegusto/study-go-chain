package chat

import (
	"net/http"

	"github.com/formegusto/study-go-chain/utils"
	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	conns = append(conns, conn)
	for {
		_, p, err := conn.ReadMessage()
		// utils.HandleErr(err)
		if err != nil {
			conn.Close()
			break
		}

		for _, aConn := range conns {
			err = aConn.WriteMessage(websocket.TextMessage, p)
			utils.HandleErr(err)
		}	
	}
}