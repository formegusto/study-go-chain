package block_chain

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data 		string
	hash 		string
	prevHash 	string 
}

func Test_1() {
	genesisBlock := block{"Genesis Block", "", ""}

	// byte slice vs string
	// for _, aByte := range "Genesis Block" {
	// 	fmt.Printf("%b\n",aByte)
	// }

	// 이는 에러 발생!
	// sha256.Sum256이 반환하는 값은 [32]byte이기 때문에 string 타입인 block.hash로는 받을 수 없다.
	// genesisBlock.hash = sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	fmt.Println(hash)
	/*
	[137 235 10 192 49 166 61 36 33 205 5 162 251 228 31 62 163 95 92 55 18 202 131 156 191 107 133 196 238 7 183 163]
	*/

	// 대부분의 hash는 16진수로 되어있다. 이 말은 base16으로 인코딩 해주면 된다는 것 이다.
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock.hash)
	/*
	89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3
	*/
}

