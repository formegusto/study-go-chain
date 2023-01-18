package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/formegusto/study-go-chain/06.persistence/db"
	"github.com/formegusto/study-go-chain/utils"
)

const DIFFICULTY int = 2

type Block struct {
	Data 		string 	`json:"data"`
	Hash 		string 	`json:"hash"`
	PrevHash 	string 	`json:"prevHash,omitempty"`
	Height		int		`json:"height"`

	// PoW; Proof of Work
	Difficulty 	int		`json:"difficulty"`
	Nonce		int		`json:"nonce"`
}

// func (b *Block) toBytes() []byte {
// 	// 1. create buffer
// 	var blockBuffer bytes.Buffer

// 	// 2. create Encoder
// 	encoder := gob.NewEncoder(&blockBuffer)

// 	// 3. run encode
// 	err := encoder.Encode(b)
// 	utils.HandleErr(err)

// 	return blockBuffer.Bytes()
// }

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)	
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)

	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		// struct -> string
		strBlock := fmt.Sprint(b)
		// string -> hash
		hash := fmt.Sprintf("%x",sha256.Sum256([]byte(strBlock)))
		fmt.Printf("Block as String:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n", strBlock, hash, target, b.Nonce)

		// mining valid
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data: data, 
		Hash: "", 
		PrevHash: prevHash, 
		Height: height,
		Difficulty: DIFFICULTY,
		Nonce: 0,
	}

	block.mine()
	// payload := block.Data + block.PrevHash + fmt.Sprintf("%d",block.Height)
	// block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()

	return &block
}