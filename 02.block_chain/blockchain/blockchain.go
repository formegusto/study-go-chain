package blockchain

import (
	"sync"
)

type Block struct {
	Data 		string 	`json:"data"`
	Hash 		string 	`json:"hash"`
	PrevHash 	string 	`json:"prevHash,omitempty"`
	Height		int		`json:"height"`
}

type blockchain struct {
}

var b *blockchain
var once sync.Once

// org. GetBlockChain
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			// b.AddBlock("Genesis Block")
		})
	}
	return b
}