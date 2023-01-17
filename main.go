package main

import (
	"github.com/formegusto/study-go-chain/02.block_chain/blockchain"
	my_cli "github.com/formegusto/study-go-chain/05.my_cli"
)

func main() {
	// persistence_test.Test()	
	blockchain.Blockchain()
	my_cli.Start()
}