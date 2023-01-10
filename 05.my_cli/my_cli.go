package my_cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/formegusto/study-go-chain/03.explorer/explorer"
	"github.com/formegusto/study-go-chain/04.rest_api/rest"
)

func usage() {
	fmt.Printf("Welcome to forme coin\n\n")
	fmt.Printf("Please use the follwing flags:\n\n")
	fmt.Printf("-port:		Set The PORT the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")

	// 프로그램 종료 함수
	os.Exit(0)
}

func Start() {
	params := os.Args
	// fmt.Println(params)
	// [/var/folders/52/4yvp7r991px0gmq4wyr894j40000gn/T/go-build874916539/b001/exe/main rest]

	// error catch
	if len(params) == 1 {
		usage()
	}

	// default value가 있어서 설정안해줘도 되긴함
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	/*
	func Parse() {
		// Ignore errors; CommandLine is set for ExitOnError.
		CommandLine.Parse(os.Args[1:])
	}
	*/
	flag.Parse()
	// fmt.Println(*port, *mode)

	switch *mode {
		case "rest":
			// start REST API
			rest.Start(*port)
		case "html":
			// start html explorer
			explorer.Start(*port)
		case "all":
			go rest.Start(*port)
			explorer.Start(*port + 1)
		default:
			usage()
	}


	// --- flagset intro code 
	// /*
	// Handler Type
	// const (
	// 	ContinueOnError ErrorHandling = iota // Return a descriptive error.
	// 	ExitOnError                          // Call os.Exit(2) or for -h/-help Exit(0).
	// 	PanicOnError                         // Call panic with a descriptive error.
	// )
	// */
	// // create flag set
	// flags := flag.NewFlagSet("flags", flag.ExitOnError)

	// // create flag item
	// // key, default value, error message 
	// // return *int 아,, 주소를 쓰는 이윸ㅋㅋ 밑에 붙음 parsed 할 때
	// portFlag := flags.Int("port", 4000, "Sets the port of the server")

	// switch params[1] {
	// 	case "explorer":
	// 		fmt.Println("Start Explorer")
	// 	case "rest":
	// 		fmt.Println("Start REST API")
	// 		// port를 찾고, port 값이 있는지, integer인지 체크하고 그에 따라 반응
	// 		flags.Parse(params[2:])
	// 	default:
	// 		welcome()
	// }

	// // parse가 제대로 되었을 때, true 반환
	// if flags.Parsed() {
	// 	fmt.Println(*portFlag)
	// 	fmt.Println("Start server")
	// }
}