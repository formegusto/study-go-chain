package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/formegusto/study-go-chain/06.persistence/db"
	"github.com/formegusto/study-go-chain/utils"
)

type Block struct {
	Data 		string 	`json:"data"`
	Hash 		string 	`json:"hash"`
	PrevHash 	string 	`json:"prevHash,omitempty"`
	Height		int		`json:"height"`
}

func (b *Block) toBytes() []byte {
	// 1. create buffer
	var blockBuffer bytes.Buffer

	// 2. create Encoder
	encoder := gob.NewEncoder(&blockBuffer)

	// 3. run encode
	err := encoder.Encode(b)
	utils.HandleErr(err)

	return blockBuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data: data, 
		Hash: "", 
		PrevHash: prevHash, 
		Height: height,
	}

	payload := block.Data + block.PrevHash + fmt.Sprintf("%d",block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()

	return &block
}