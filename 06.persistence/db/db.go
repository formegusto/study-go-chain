// main.go 에서 직접 호출할 일 없음
package db

import (
	"fmt"

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

// key : DB 에서 block 회수 -> hash
// value : byte 형태의 data, block 전체를 byte 형태로 변환
// boltdb byte 형태의 data만 받는다.
func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(tx *bolt.Tx) error {
		// bucket에 저장작업
		// 1. bucket read
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)

		return err
	})
	utils.HandleErr(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		// 1. bucket read
		bucket := tx.Bucket([]byte(dataBucket))

		// 2. newestHash 와 height 정보가 들어갈 거라!
		err := bucket.Put([]byte("checkpoint"), data)

		return err
	})
	utils.HandleErr(err)
}