package blockchain

import (
	"fmt"
	"sync"

	"github.com/formegusto/study-go-chain/06.persistence/db"
	"github.com/formegusto/study-go-chain/utils"
)

const (
	defaultDifficulty 	int = 2
	difficultyInterval 	int = 5
	// 2분 간격으로 생성되었으면 좋겠다.
	blockInterval 		int = 2

    allowedRange		int = 2
)

type blockchain struct {
	// cursor or pointer 가장 최신의 블록 정보를 기록해놓는다.
	NewestHash 			string 	`json:"newestHash"`
	Height				int		`json:"height"`

	// difficulty pointer, 이전 블록의 difficulty (가장 최신의)
	CurrentDifficulty	int		`json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	fmt.Println("Chain Restoring...")
	utils.FromBytes(b, data)
}

func (b* blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock() {
	// prevHash와 height는
	// 블록체인에서 알아야 한다. 
	// 가장 최신의 블록을
	// block := Block{data,"", b.NewestHash, b.Height + 1}
	// 1. DB 관련 코드가 많아서 바꿈
	block := createBlock(b.NewestHash, b.Height + 1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
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

func (b *blockchain) recalculateDifficulty() int {
	// 결론적으로 우리는
	// difficultyInterval 개를 만드는 데에 
	// X blockInterval 분만큼의 시간을 기대한다.
	// 예상 기대 시간 = 10분
	allBlocks := b.Blocks()

	// chain.Blocks()의 코드를 보면
	// 가장 최신의 블록부터(NewestHash) 조회하는 것을 알 수 있다.
	// 그래서 0번 인덱스가 가장 최신의 블록
	newestBlock := allBlocks[0]

	// 가장 최근에 Difficulty가 재설정된 블록을 찾아야 한다.
	// 5개 단위로 재설정하기 때문에 difficultyInterval을 사용한다.
	lastRecalculatedBlock := allBlocks[difficultyInterval - 1]

	// 두 블록 사이에 걸린 시간
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)

	// 우리가 말하는 예상 기대 시간
	expectedTime := difficultyInterval * blockInterval

	// 10분이 아니면, 무조건 올리거나 내리거나 하는 엄격한 면이 있음
	// 그래서 범위를 적용한다. (allowedRange)
	// fmt.Printf("actualTime:%d\nexpectedTime:%d\n\n", actualTime, expectedTime)
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	// Default Difficulty is 0
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height % difficultyInterval == 0 {
		// recalculate the difficulty
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

// 우리는 이제 uTxOuts만 가지고 오면 된다. 아직 사용되지 않은 transaction out!
// func (b *blockchain) txOuts() []*TxOut {
// 	var txOuts []*TxOut
// 	blocks := b.Blocks()
// 	// blocks search
// 	for _, block := range blocks {
// 		// transactions search
// 		for _, tx := range block.Txs {
// 			txOuts = append(txOuts, tx.TxOuts...)
// 		}
// 	}

// 	return txOuts
// }
// func (b *blockchain) TxOutsByAddress(address string) []*TxOut {
// 	var ownedTxOuts []*TxOut
// 	txOuts := b.txOuts()

// 	for _, txOut := range txOuts {
// 		if txOut.Owner == address {
// 			ownedTxOuts = append(ownedTxOuts, txOut)
// 		}
// 	}
// 	return ownedTxOuts
// }

// 아직 input에서 사용되지 않은 output을 넘겨줄 것 이다.
func (b *blockchain) UTxOutsByAddress(address string) []*UTxOut {
	// input은 이제 output을 찾기 위한 표지판 역할이다.
	// 1. TxID를 찾아서,,
	var uTxOuts []*UTxOut
	// spend transaction outputs
	creatorTxs := make(map[string]bool)

	for _, block := range b.Blocks() {
		for _, tx := range block.Txs {
			// 모든 input은 어떠한 output으로 부터 나왔다.
			for _, input := range tx.TxIns {
				if input.Owner == address {
					// 사용하는 output을 찾아낼 수 있다.
					// map 처럼 저장해야 한다.
					creatorTxs[input.TxID] = true
				}
			}
			
			// 여기서 확인함
			for index, output := range tx.TxOuts {
				if output.Owner == address {
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


func (b* blockchain) BalanceByAddress(address string) int {
	txOuts := b.UTxOutsByAddress(address)
	var amount int
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

// org. GetBlockChain
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				Height: 0,
			}
			fmt.Printf("NewestHash: %s\nHeight: %d\n\n", b.NewestHash, b.Height)

			// search for checkpoint on the db
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock()
			} else {
				// restore b from bytes
				b.restore(checkpoint)
			}
		})
	}

	fmt.Printf("NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
	return b
}