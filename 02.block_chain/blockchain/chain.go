package blockchain

import (
	"sync"
)

type blockchain struct {
	// cursor or pointer 가장 최신의 블록 정보를 기록해놓는다.
	NewestHash 	string 	`json:"newestHash"`
	Height		int		`json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	// prevHash와 height는
	// 블록체인에서 알아야 한다. 
	// 가장 최신의 블록을
	// block := Block{data,"", b.NewestHash, b.Height + 1}
	// 1. DB 관련 코드가 많아서 바꿈
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

// org. GetBlockChain
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}