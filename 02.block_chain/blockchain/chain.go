package blockchain

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/formegusto/study-go-chain/06.persistence/db"
	"github.com/formegusto/study-go-chain/utils"
)

const (
	defaultDifficulty 	int = 2
	difficultyInterval 	int = 5
	blockInterval 		int = 2
    allowedRange		int = 2
)

type blockchain struct {
	NewestHash 			string 	`json:"newestHash"`
	Height				int		`json:"height"`
	CurrentDifficulty	int		`json:"currentDifficulty"`
	m					sync.Mutex
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func persistBlockchain(b* blockchain)  {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height + 1, getDifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockchain(b)
}

func Blocks(b *blockchain) []*Block {
	b.m.Lock()
	defer b.m.Unlock()
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func Txs(b *blockchain) []*Tx {
	var txs []*Tx
	for _, block := range Blocks(b) {
		txs = append(txs, block.Txs...)
	}
	return txs
}

func FindTx(b *blockchain, targetId string) *Tx{
	for _, tx := range Txs(b) {
		if tx.Id == targetId {
			return tx
		}
	}
	return nil
}

func recalculateDifficulty(b *blockchain)  int {
	allBlocks := Blocks(b)
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval - 1]

	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval

	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}

	return b.CurrentDifficulty
}

func getDifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height % difficultyInterval == 0 {
		return recalculateDifficulty(b)
	} else {
		return b.CurrentDifficulty
	}
}

func UTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	var uTxOuts []*UTxOut
	creatorTxs := make(map[string]bool)

	for _, block := range Blocks(b) {
		for _, tx := range block.Txs {
			for _, input := range tx.TxIns {
				if input.Signature == "COINBASE" {
					break
				}
				if FindTx(b, input.TxID).TxOuts[input.Index].Address == address {
					creatorTxs[input.TxID] = true
				}
			}

			for index, output := range tx.TxOuts {
				if output.Address == address {
					if _, ok := creatorTxs[tx.Id]; !ok {
						uTxOut := &UTxOut{tx.Id, index, output.Amount}
						if !isOnMempool(uTxOut) {
							uTxOuts = append(uTxOuts, uTxOut)
						}
					}
				}	
			}
		}
	}

	return uTxOuts
}

func BalanceByAddress(address string, b* blockchain) int {
	txOuts := UTxOutsByAddress(address, b)
	var amount int
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

func Blockchain() *blockchain {
	once.Do(func() {
		b = &blockchain{
			Height: 0,
		}
		// fmt.Printf("NewestHash: %s\nHeight: %d\n\n", b.NewestHash, b.Height)

		checkpoint := db.Checkpoint()
		if checkpoint == nil {
			b.AddBlock()
		} else {
			b.restore(checkpoint)
		}
	})

	// fmt.Printf("NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
	return b
}

func (b *blockchain) Replace(newBlocks []*Block) {
	b.m.Lock()
	defer b.m.Unlock()
	b.CurrentDifficulty = newBlocks[0].Difficulty
	b.Height = len(newBlocks)
	b.NewestHash = newBlocks[0].Hash
	persistBlockchain(b)

	db.EmptyBlocks()
	for _, block := range newBlocks {
		persistBlock(block)
	}
}

func Status(b *blockchain, rw http.ResponseWriter) {
	b.m.Lock()
	defer b.m.Unlock()

	json.NewEncoder(rw).Encode(b)
}