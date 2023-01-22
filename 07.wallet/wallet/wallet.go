package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"
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

func encodeBigInts(a, b *big.Int) string {
	bytes := append(a.Bytes(), b.Bytes()...)
	return fmt.Sprintf("%x", bytes)
}

func aFromK(key *ecdsa.PrivateKey) string {
	x := key.X
	y := key.Y
	z := encodeBigInts(x, y)

	return z
}

func sign(payload string, w wallet) string {
	payloadAsBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsBytes)
	utils.HandleErr(err)

	signature := encodeBigInts(r, s)
	return signature
}

func restoreBigInt(payload string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(payload)
	if err != nil {
		return nil, nil, err
	}

	firstHalfBytes, secondHalfBytes := bytes[:len(bytes) / 2], bytes[len(bytes) / 2:]
	bigR, bigS := &big.Int{}, &big.Int{}
	bigR.SetBytes(firstHalfBytes)
	bigS.SetBytes(secondHalfBytes)

	return bigR, bigS, nil
}

func verify(signature, payload, address string) bool {
	// 1. restore signature
	r, s, err := restoreBigInt(signature)
	utils.HandleErr(err)

	// 2. restore publicKey
	x, y, err := restoreBigInt(address)
	utils.HandleErr(err)
	publicKey := ecdsa.PublicKey {
		Curve: 	elliptic.P256(),
		X: 		x,
		Y: 		y,
	}

	// 3. Verify!
	payloadAsBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	ok := ecdsa.Verify(&publicKey, payloadAsBytes, r, s)

	return ok
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