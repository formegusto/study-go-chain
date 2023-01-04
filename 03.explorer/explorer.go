package explorer_test

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

// 2개의 인자를 받는다.
// 1. http.ResponseWriter : 유저에게 보내고 싶은 데이터 정의
// 2. *http.Request : request로는 큰 파일이 올 수도 있기 때문에 복사보다는 실제 값을 이용하는 것을 지향한다.
func home(rw http.ResponseWriter, r *http.Request) {
	// console에 출력이 아닌, Writer(출력 스트림)에 출력하는 기능
	fmt.Fprint(rw, "Hello from home!")
}

func Open() {
	// routes
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)

	// 에러가 발생하면 error 출력 후 os.Exit(1) 으로 프로그램 종료
	log.Fatal(http.ListenAndServe(port, nil))
}