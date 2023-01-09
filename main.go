package main

import (
	explorer_test "github.com/formegusto/study-go-chain/03.explorer"
	rest_test "github.com/formegusto/study-go-chain/04.rest_api"
)

func main() {
	go explorer_test.Open()
	rest_test.Open()
}