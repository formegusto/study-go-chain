package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/formegusto/study-go-chain/utils"
)

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	// 1. message 
	message := "I love you"

	// 2. hash message
	hashedMessage := utils.Hash(message)

	// 3. sign hash
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	// 4. what?
	utils.HandleErr(err)
	fmt.Printf("R:%d\nS:%d\n", r,s)

}