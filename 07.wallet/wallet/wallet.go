package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	"github.com/formegusto/study-go-chain/utils"
)

const (
	filename string = "wallets/formecoin.wallet"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
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

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}

		// 1. has a wallet already?
		isHas := hasWalletFile()
		if isHas {
			// 2-a. yes -> restore from file
		} else {
			// 2-b. no -> create prv key, save to file
			key := createPrvKey()
			persistKey(key)
			w.privateKey = key
		}
	}

	return w
}