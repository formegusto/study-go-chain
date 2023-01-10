package my_cli

import (
	"flag"
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

	/*
	Handler Type
	const (
		ContinueOnError ErrorHandling = iota // Return a descriptive error.
		ExitOnError                          // Call os.Exit(2) or for -h/-help Exit(0).
		PanicOnError                         // Call panic with a descriptive error.
	)
	*/
	// create flag set
	flags := flag.NewFlagSet("flags", flag.ExitOnError)

	// create flag item
	// key, default value, error message 
	// return *int 아,, 주소를 쓰는 이윸ㅋㅋ 밑에 붙음 parsed 할 때
	portFlag := flags.Int("port", 4000, "Sets the port of the server")

	switch params[1] {
		case "explorer":
			fmt.Println("Start Explorer")
		case "rest":
			fmt.Println("Start REST API")
			// port를 찾고, port 값이 있는지, integer인지 체크하고 그에 따라 반응
			flags.Parse(params[2:])
		default:
			welcome()
	}

	// parse가 제대로 되었을 때, true 반환
	if flags.Parsed() {
		fmt.Println(*portFlag)
		fmt.Println("Start server")
	}
}