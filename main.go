package main

import (
	my_cli "github.com/formegusto/study-go-chain/05.my_cli"
	"github.com/formegusto/study-go-chain/06.persistence/db"
)

func main() {
	defer db.Close()
	my_cli.Start()
	// wallet_test.Test()
	// p2p_test.BasicTest()
}