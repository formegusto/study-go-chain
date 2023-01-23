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
	go countToTen(c)
	receive(c)

	// go CountToTen("two")
	// *. channel 없이 기다리는 방법
	// for {}
}

func receive(c <-chan int) {
	for {
		a, ok := <-c
		if !ok {
			fmt.Println("We are done.")
			break
		}
		fmt.Printf("received %d\n", a)
	}
}

func countToTen(c chan<- int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i
		// fmt.Println(i)
	}
	close(c)
}