package p2p_test

import (
	"fmt"
	"time"
)

/*
go routine basic
*/
func BasicTest() {
	go CountToTen("one")
	go CountToTen("two")
	for {}
}

func CountToTen(name string) {
	for i := range [10]int{} {
		fmt.Println(name,i)
		time.Sleep(1 * time.Second)
	}
}