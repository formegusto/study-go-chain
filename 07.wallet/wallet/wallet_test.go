package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"testing"
)

const (
	testKey 		string = "307702010104205605a3b9bb33a03c2d688e63df6549227875a9c4dc164f3cd11e28a86f837938a00a06082a8648ce3d030107a1440342000461aa502eae682a24d8ddbdbcb464dccd4e0547777d0f2a6fdb22a356b078c2a8957540500f85df76497dc4b1bc705679995b6692f4a76aaa1ceb697ea42dc262"
	// hex 값의 payload가 필요함
	testPayload 	string = "000f5dd086ba6d30d1d9e511847acaab138ad1a002fe105d16472f2ebeec3c07"
	testSig 		string = "5377ffff729ffaf2f2258bbb4141f85e490138fda1f99ae63fc87efa16a45db034cd4aa921c32977bf9f580446157cd1a0d39d64e07276ec398b9881764209db"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)

	w.privateKey = key
	w.Address = aFromK(key)

	return w
}

func TestSign(t *testing.T) {
	signature := Sign(testPayload, *makeTestWallet())
	// 서명은 hash와 달리 랜덤하다.
	// 그래서 값이 일치하는지 확인할 수 없다.
	// 그래서 hexDecimal여부를 확인한다.
	_, err := hex.DecodeString(signature)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s\n", signature)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input	string
		ok		bool
	}
	tests := []test{
		{
			input: testPayload,
			ok: true,
		},
		{
			input: "040f5dd086ba6d30d1d9e511847acaab138ad1a002fe105d16472f2ebeec3c07",
			ok: false,
		},
	}

	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
	
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInt("xx")
	if err == nil {
		t.Error("restoreBigInt() should return error when payload is not hex.")
	}
}