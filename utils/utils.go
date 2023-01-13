package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error)  {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(i)
	HandleErr(err)
	return buffer.Bytes()
}