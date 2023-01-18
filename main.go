package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hash := sha256.Sum256([]byte("hello"))
	fmt.Printf("%x\n",hash)

	// 1차 시도
	hash = sha256.Sum256([]byte("hello0"))
	fmt.Printf("%x\n",hash)

	// 2차 시도
	hash = sha256.Sum256([]byte("hello1"))
	fmt.Printf("%x\n",hash)

}