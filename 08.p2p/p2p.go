package p2p_test

import (
	"fmt"
	"time"
)

/*
go routine basic
*/
func BasicTest() {
	c := make(chan int)
	go CountToTen(c)

	fmt.Println("Blocking")
	for {
		a := <-c
		fmt.Printf("received %d\n", a)
	}

	// go CountToTen("two")
	// *. channel 없이 기다리는 방법
	// for {}
}

func CountToTen(c chan int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i
		// fmt.Println(i)
	}
}