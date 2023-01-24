package p2p

import (
	"encoding/json"
	"fmt"

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

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind: kind,
		Payload: utils.ToJSON(payload),
	}

	return utils.ToJSON(m)
}

func sendNewestBlock(p *peer) {
	// 최신의 블록 받아오기
	fmt.Printf("Sending newest block to %s\n", p.key)
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)

	// 변환해서 보내기
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func requestAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksRequest, nil)
	p.inbox <- m
}

func sendAllBlocks(p *peer) {
	m := makeMessage(MessageAllBlocksResponse, blockchain.Blocks(blockchain.Blockchain()))
	p.inbox <- m
}

func handleMessage(m *Message, p *peer) {
	// fmt.Printf("Peer: %s, Sent a message with kind of: %d\n", p.key, m.Kind)
	switch m.Kind {
		case MessageNewestBlock:
			fmt.Printf("Received the newest block from %s\n", p.key)
			var payload blockchain.Block
			err := json.Unmarshal(m.Payload, &payload)
			utils.HandleErr(err)

			b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
			utils.HandleErr(err)
			if payload.Height >= b.Height {
				// request all the blocks from 4000
				fmt.Printf("Requesting all blocks from %s\n", p.key)
				requestAllBlocks(p)
			} else {
				// send 4000 our block
				fmt.Printf("Sending newest block to %s\n", p.key)
				sendNewestBlock(p)
			}
		case MessageAllBlocksRequest:
			fmt.Printf("%s wants all the blocks\n", p.key)
			sendAllBlocks(p)
		case MessageAllBlocksResponse:
			fmt.Printf("Received all the blocks from %s\n", p.key)
			var payload []*blockchain.Block
			err := json.Unmarshal(m.Payload, &payload)
			utils.HandleErr(err)

			blockchain.Blockchain().Replace(payload)

	}
}