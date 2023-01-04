package explorer_test

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/formegusto/study-go-chain/02.block_chain/blockchain"
)

const (
	port string = ":4000"
	templateDir string = "templates/"
)
var templates *template.Template

type homeData struct {
	// template에서 공우되는 것 마저,, 대문자로 공유해준다!
	PageTitle string
	Blocks []*blockchain.Block
}

// 2개의 인자를 받는다.
// 1. http.ResponseWriter : 유저에게 보내고 싶은 데이터 정의
// 2. *http.Request : request로는 큰 파일이 올 수도 있기 때문에 복사보다는 실제 값을 이용하는 것을 지향한다.
func home(rw http.ResponseWriter, r *http.Request) {
	// console에 출력이 아닌, Writer(출력 스트림)에 출력하는 기능
	// fmt.Fprint(rw, "Hello from home!")
	
	// * templates
	// tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Must Function을 사용해주면 위에서 작업한 error를 모듈에서 알아서 처리해준다.

	// Writer, data 를 요구한다.
	// Template 에서 받고자 하는 데이터를 정의해준다.
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home",data)
}

func Open() {
	// template preload
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	// templates 를 쓰는 것이 중요하다. 빈 곳에 저장 후 추가 저장한다는 느낌임
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	// routes
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)

	// 에러가 발생하면 error 출력 후 os.Exit(1) 으로 프로그램 종료
	log.Fatal(http.ListenAndServe(port, nil))
}