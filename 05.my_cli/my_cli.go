package my_cli

import (
	"fmt"
	"os"
)

func welcome() {
	fmt.Printf("Welcome to forme coin\n\n")
	fmt.Printf("Please use the follwing commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:		Start the REST API (recommended)\n\n")

	// 프로그램 종료 함수
	os.Exit(0)
}

func Test() {
	params := os.Args
	// fmt.Println(params)
	// [/var/folders/52/4yvp7r991px0gmq4wyr894j40000gn/T/go-build874916539/b001/exe/main rest]

	// error catch
	if len(params) < 2 {
		welcome()
	}

	// option은 3번째 부터 붙는 걸로!
	fmt.Println(params[2:])

	switch params[1] {
		case "explorer":
			fmt.Println("Start Explorer")
		case "rest":
			fmt.Println("Start REST API")
		default:
			welcome()
	}
}