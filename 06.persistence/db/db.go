// main.go 에서 직접 호출할 일 없음
package db

import (
	"github.com/boltdb/bolt"
	"github.com/formegusto/study-go-chain/utils"
)

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
)
var db *bolt.DB

// db initialize
func DB() *bolt.DB{
	if db == nil {
		// init db 
		// path는 database 이름이고, 파일이 없으면 자동으로 생성해준다.
		// 1. path, 2. permision, 3. options
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		utils.HandleErr(err)
		db = dbPointer

		// Update Transaction
		// Write and Read
		// bucket 아니면 error를 리턴
		err = db.Update(func(tx *bolt.Tx) error {
			// 우리는 bucket이 필요한게 아니고, 생성해야 해서 그냥 _ 처리
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))

			return err
		})
		utils.HandleErr(err)
	}
	return db
}
