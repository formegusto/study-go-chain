package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/formegusto/study-go-chain/utils"
)

const (
	filename string = "wallets/formecoin.wallet"
)

type wallet struct {
	privateKey 	*ecdsa.PrivateKey
	Address		string
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat(filename)
	return os.IsExist(err)
}

func createPrvKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)

	err = os.WriteFile(filename, bytes, 0644)
	utils.HandleErr(err)
}

func restoreKey() (privateKey *ecdsa.PrivateKey) {
	keyAsBytes, err := os.ReadFile(filename)
	utils.HandleErr(err)
	privateKey, err = x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleErr(err)
	return
}

func aFromK(key *ecdsa.PrivateKey) string {
	x := key.X.Bytes()
	y := key.Y.Bytes()
	z := append(x, y...)

	return fmt.Sprintf("%x", z)
}

func sign(payload string, w *wallet) string {
	payloadAsBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsBytes)
	utils.HandleErr(err)

	signature := append(r.Bytes(), s.Bytes()...)
	return fmt.Sprintf("%x", signature)
}

func verify(signature, payload, publicKey string) bool {
	return true
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}

		// 1. has a wallet already?
		isHas := hasWalletFile()
		if isHas {
			// 2-a. yes -> restore from file
			w.privateKey = restoreKey()
		} else {
			// 2-b. no -> create prv key, save to file
			key := createPrvKey()
			persistKey(key)
			w.privateKey = key
		}
		w.Address = aFromK(w.privateKey)
	}

	return w
}