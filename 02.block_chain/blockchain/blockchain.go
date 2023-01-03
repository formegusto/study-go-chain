package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data 		string
	hash 		string
	prevHash 	string 
}

type blockchain struct {
	blocks []*block
}

// 1. singletone variable
var b *blockchain
var once sync.Once

// 2. control function
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x",hash)
}

func getLastHash() string {
	totalBlock := len(GetBlockchain().blocks)
	if totalBlock == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlock - 1].hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()

	return &newBlock
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.blocks = append(b.blocks, createBlock("GenesisBlock"))
		})
	}
	return b
}


