package blockchain

import (
	"errors"
	"time"

	"github.com/formegusto/study-go-chain/utils"
)

const (
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx 
}

// Blockchain의 b변수와는 다르게
// 초기화가 필요없다. 데이터베이스에 넣지 않을 것 이기 때문이다.
// 왜냐하면 transaction이 성립되면 mempool에서 transaction은 없어지고, 블록 내부로 들어와
// 블록체인 DB에 상주하게 된다.
// mempool은 메모리에서만 존재한다.
var Mempool *mempool = &mempool{}

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

func makeTx(from, to string, amount int) (*Tx, error) {
	// 누군가가 얼마만큼 가지고 있는지 알고 싶다면 출력값을 참고해야 한다.
	// transaction을 시작하고 싶다면 입력값을 만들면 된다.
	// 그리고 입력값은 blockchain 가지고 있는 나의 돈 이다.

	// 이 말은 blockchain이 아니라, 인간이 만드는 transaction을 발생시키고 싶다면,
	// 인간이 트랜잭션의 출력값을 가지고 있어야 하고, 이 출력값을 트랜잭션 입력값으로 다시 변경해주어야 한다는 것을 말한다.
	// transaction input은 예전의 transaction output 이다.
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("Not Enough money")
	}

	// 여러개의 TxOut 중에서 사용자가 주어야 하는 TxOut 만큼만 사용하면 된다.
	// 목표는 크거나 같을 때 까지,
	var txIns []*TxIn
	var txOuts []*TxOut
	total := 0
	oldTxOuts := Blockchain().TxOutsByAddress(from)
	for _, txOut := range oldTxOuts {
		if total > amount {
			break
		}
		txIn := &TxIn{txOut.Owner, txOut.Amount}
		txIns = append(txIns, txIn)
		total += txIn.Amount
	}

	// 잔돈이 있을 수 있기 때문에
	change := total - amount
	// 잔돈용 Tx 생성
	if change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	// amount transaction output
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)

	tx := &Tx{
		Id:"",
		Timestamp: int(time.Now().Unix()),
		TxIns: txIns,
		TxOuts: txOuts,
	}
	tx.getId()

	return tx, nil
}

// 여기서 보내는 사람(from)은 필요하지 않다.
// 보내는 사람은 function이 아닌, wallet에서 받아올 것이기 때문이다.
func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("forme",to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}