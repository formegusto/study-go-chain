// utils contains function to be used across the application.
package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var logFn = log.Panic

func HandleErr(err error)  {
	if err != nil {
		logFn(err)
	}
}

func ToBytes(i interface{}) []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(i)
	HandleErr(err)
	return buffer.Bytes()
}

// FromBytes takses an interface and data and then will encode the data to the interface
func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(i)
	HandleErr(err)	
}

// Hash takes an interface, hashes it and returns the hex encoding of the hash.
func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func Splitter(s string, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r) - 1 < i {
		return ""
	}
	return r[i]
}

func ToJSON(i interface{}) []byte {
	r, err := json.Marshal(i)
	HandleErr(err)

	return r
}