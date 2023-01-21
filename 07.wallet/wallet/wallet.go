package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/formegusto/study-go-chain/utils"
)

const (
	hashedMessage string = "c33084feaa65adbbbebd0c9bf292a26ffc6dea97b170d88e501ab4865591aafd"
)

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	// 1. message 
	message := "I love you"

	// 2. hash message
	hashedMessage := utils.Hash(message)
	fmt.Println(hashedMessage)

	// 3. sign hash
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	// 4. what?
	utils.HandleErr(err)
	// fmt.Printf("R:%d\nS:%d\n", r,s)

	// 5. verify
	ok := ecdsa.Verify(&privateKey.PublicKey, hashAsBytes, r, s)
	fmt.Println(ok)
}