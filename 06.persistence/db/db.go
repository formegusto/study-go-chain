// main.go 에서 직접 호출할 일 없음
package db

import (
	"fmt"
	"os"

	"github.com/formegusto/study-go-chain/utils"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName = "blockchain"
	dataBucket = "data"
	blocksBucket = "blocks"

	checkpoint = "checkpoint"
)
var db *bolt.DB

func getDBName() string{
	port := os.Args[2][6:]
	return fmt.Sprintf("%s_%s.db", dbName, port)
}

// db initialize
func DB() *bolt.DB{
	if db == nil {
		// init db 
		// path는 database 이름이고, 파일이 없으면 자동으로 생성해준다.
		// 1. path, 2. permision, 3. options
		dbPointer, err := bolt.Open(getDBName(), 0600, nil)
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

func Close() {
	DB().Close()
}

// key : DB 에서 block 회수 -> hash
// value : byte 형태의 data, block 전체를 byte 형태로 변환
// boltdb byte 형태의 data만 받는다.
func SaveBlock(hash string, data []byte) {
	// fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(tx *bolt.Tx) error {
		// bucket에 저장작업
		// 1. bucket read
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)

		return err
	})
	utils.HandleErr(err)
}

func SaveCheckpoint(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		// 1. bucket read
		bucket := tx.Bucket([]byte(dataBucket))

		// 2. newestHash 와 height 정보가 들어갈 거라!
		err := bucket.Put([]byte(checkpoint), data)

		return err
	})
	utils.HandleErr(err)
}

func Checkpoint()[]byte {
	var data []byte
	// only read transaction
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))

		return nil
	})

	return data
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))

		return nil
	})

	// nil or byte data
	return data
}

func EmptyBlocks() {
	DB().Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(blocksBucket))
		utils.HandleErr(err)

		_, err = tx.CreateBucket([]byte(blocksBucket))
		utils.HandleErr(err)

		return nil
	})
}