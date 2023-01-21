package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/formegusto/study-go-chain/utils"
)

const (
	privateKey string = "3077020101042045ca6857700ace5906b3bea921d69992ceff772485aed28537953750060c5b25a00a06082a8648ce3d030107a14403420004a2c6839013e9b09f4e68f7765531feba60d11ed384d119f36594173976d9fd33d802da0de74ad5ea91b5e614bac048721a0db4039b5590fab78243d619b2241d"
	hashedMessage string = "c33084feaa65adbbbebd0c9bf292a26ffc6dea97b170d88e501ab4865591aafd"
	signature string = "d61fdba987524bb6b345028bf456e9dabebd7ebf889f885d6f78ff36e751b9ce47ad557e186ae01232e8cca0218c4c845e7c572fbcfa1e43c1f265a1ddde1c3d"
)

func Start() {
	// 1. privateKey -> privateBytes
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	restoredKey, err = x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	// 2. signature slice
	sigBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)
	rBytes, sBytes := sigBytes[:len(sigBytes) / 2], sigBytes[len(sigBytes) / 2:]
	fmt.Println(rBytes, sBytes)

	// 3. bigR, bigS restore
	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
	fmt.Println(bigR, bigS)
}

func BasicProcess() {
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