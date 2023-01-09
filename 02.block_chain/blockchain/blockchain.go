package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data 		string 	`json:"data"`
	Hash 		string 	`json:"hash"`
	PrevHash 	string 	`json:"prevHash,omitempty"`
	Height		int		`json:"height"`
}

type blockchain struct {
	blocks []*Block
}

// 1. singletone variable
var b *blockchain
var once sync.Once

// 2. control function
func (b *Block) calculateHash() {
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

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockchain().blocks) + 1}
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

func (b blockchain) AllBlocks() []*Block {
	return b.blocks
}

func (b *blockchain) GetBlock(height int) *Block{
	// height는 1부터
	// array는 0부터
	return b.blocks[height - 1]
}


