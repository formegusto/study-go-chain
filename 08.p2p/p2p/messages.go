package p2p

import (
	"encoding/json"

	"github.com/formegusto/study-go-chain/02.block_chain/blockchain"
	"github.com/formegusto/study-go-chain/utils"
)

type MessageKind int

const (
	MessageNewestBlock			MessageKind	= iota 
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

type Message struct {
	Kind		MessageKind
	Payload		[]byte
}

func (m *Message) addPayload(p interface{}) {
	b, err := json.Marshal(p)
	utils.HandleErr(err)

	m.Payload = b
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind: kind,
	}
	m.addPayload(payload)

	mJson, err := json.Marshal(m)
	utils.HandleErr(err)
	
	return mJson
}

func sendNewestBlock(p *peer) {
	// 최신의 블록 받아오기
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)

	// 변환해서 보내기
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}