package p2p

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type peers struct {
	v	map[string]*peer
	m 	sync.Mutex
}

// var Peers map[string]*peer = make(map[string]*peer)
var Peers peers = peers{
	v: make(map[string]*peer),
}

type peer struct {
	key		string
	address	string
	port	string
	conn 	*websocket.Conn
	inbox	chan	[]byte
}

func AllPeers(p *peers) []string {
	p.m.Lock()
	defer p.m.Unlock()

	var peers []string

	for key := range p.v {
		peers = append(peers, key)
	}

	return peers
}

func (p *peer) close() {
	// 어느 곳에서나 잠겨있는 함수, 오로지 실행 가능은 이 함수
	Peers.m.Lock()
	// 잠금해제
	defer Peers.m.Unlock()
	p.conn.Close()
	delete(Peers.v, p.key)
}

func (p *peer) read() {
	defer p.close()
	// delete peer in case of error
	for {
		// _, m, err := p.conn.ReadMessage()
		m := Message{}
		err := p.conn.ReadJSON(&m)
		if err != nil {
			break
		}
		handleMessage(&m, p)
	}
}

func (p *peer) write() {
	defer p.close()
	for {
		m, ok := <- p.inbox
		if !ok {
			break
		}
		p.conn.WriteMessage(websocket.TextMessage, m)
	}
}

func initPeer(conn *websocket.Conn, address ,port string) *peer{
	Peers.m.Lock()
	defer Peers.m.Unlock()
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		key:		key,
		address:	address,
		port:		port,
		conn:		conn,
		inbox:		make(chan []byte),
	}
	go p.read()
	go p.write()
	Peers.v[key] = p
	return p
}