package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data 		string
	Hash 		string
	PrevHash 	string 
}

type blockchain struct {
	blocks []*block
}

// 1. singletone variable
var b *blockchain
var once sync.Once

// 2. control function
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x",hash)
}

func getLastHash() string {
	totalBlock := len(GetBlockchain().blocks)
	if totalBlock == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlock - 1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()

	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

func (b blockchain) AllBlocks() []*block {
	return b.blocks
}


