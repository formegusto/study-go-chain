package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	difficulty := 6
	target := strings.Repeat("0", difficulty)
	nonce := 1
	data := "hello"

	// 찾기
	for {
		bHash := sha256.Sum256([]byte(data + fmt.Sprint(nonce)))
		hash := fmt.Sprintf("%x", bHash)
		fmt.Printf("Hash:%s\nTarget:%s\nNonce:%d\n\n", hash, target, nonce)

		if strings.HasPrefix(hash, target) {
			break
		} else {
			nonce++
		}
	}
}