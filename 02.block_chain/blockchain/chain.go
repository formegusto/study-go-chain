package blockchain

import (
	"fmt"
	"sync"

	"github.com/formegusto/study-go-chain/06.persistence/db"
	"github.com/formegusto/study-go-chain/utils"
)

type blockchain struct {
	// cursor or pointer 가장 최신의 블록 정보를 기록해놓는다.
	NewestHash 	string 	`json:"newestHash"`
	Height		int		`json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	fmt.Println("Chain Restoring...")
	utils.FromBytes(b, data)
}

func (b* blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	// prevHash와 height는
	// 블록체인에서 알아야 한다. 
	// 가장 최신의 블록을
	// block := Block{data,"", b.NewestHash, b.Height + 1}
	// 1. DB 관련 코드가 많아서 바꿈
	block := createBlock(data, b.NewestHash, b.Height + 1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func(b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			// Genesis block 까지 탐색
			break
		}
	}
	return blocks
}

// org. GetBlockChain
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			fmt.Printf("NewestHash: %s\nHeight: %d\n\n", b.NewestHash, b.Height)

			// search for checkpoint on the db
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				// restore b from bytes
				b.restore(checkpoint)
			}
		})
	}

	fmt.Printf("NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
	return b
}