package blockchain

import (
	"time"

	"github.com/formegusto/study-go-chain/utils"
)

const (
	minerReward int = 50
)

type Tx struct {
	Id 			string		`json:"id"`
	Timestamp 	int			`json:"timestamp"`
	TxIns 		[]*TxIn		`json:"txIns"`
	TxOuts 		[]*TxOut	`json:"txOuts"`
}
func (tx *Tx) getId() {
	tx.Id = utils.Hash(tx)
}

type TxIn struct {
	Owner 	string
	Amount 	int
}

type TxOut struct {
	Owner 	string
	Amount 	int
}

// address is miner address
func makeCoinbaseTx(address string) *Tx {
	// owner : 소유주
	// amount : 채굴자에게 지급할 액수의 수량
	txIns := []*TxIn {
		{"COINBASE", minerReward},
	}
	// owner : 채굴자의 주소
	// amount : 거래 총량
	txOuts := []*TxOut {
		{address, minerReward},
	}
	tx := Tx{
		Id: "",
		Timestamp: int(time.Now().Unix()),
		TxIns: txIns,
		TxOuts: txOuts,
	}
	tx.getId()
	return &tx
}