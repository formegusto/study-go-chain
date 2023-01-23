package p2p

import (
	"fmt"
	"net/http"
	"time"

	"github.com/formegusto/study-go-chain/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)

	for {
		_, p, err := conn.ReadMessage()
		// utils.HandleErr(err)
		if err != nil {
			conn.Close()
			break
		}
		fmt.Printf("Just got: %s\n\n", p)
		
		time.Sleep(5 * time.Second)

		message := fmt.Sprintf("we also think that: %s", p)
		err = conn.WriteMessage(websocket.TextMessage, []byte(message))
		utils.HandleErr(err)
	}
}