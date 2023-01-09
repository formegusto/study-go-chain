package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/formegusto/study-go-chain/02.block_chain/blockchain"
	"github.com/formegusto/study-go-chain/utils"
	"github.com/gorilla/mux"
)

var port string

type url string
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s",port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL 		url    `json:"url"`
	Method 		string `json:"method"`
	Description string `json:"description"`
	Payload		string `json:"payload,omitempty"`
	AdminMsg	string `json:"-"`
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription {
		{
			URL: 			url("/"),
			Method: 		"GET",
			Description: 	"See Documentation",
			AdminMsg: 		"This is hide field",
		},
		{
			URL: 			url("/blocks"),
			Method: 		"GET",
			Description: 	"See All Blocks",
			AdminMsg: 		"This is hide field",
		},
		{
			URL: 			url("/blocks"),
			Method: 		"POST",
			Description: 	"Add A Block",
			Payload: 		"data:string",
			AdminMsg: 		"This is hide field",
		},
		{
			URL: 			url("/blocks/{height}"),
			Method: 		"GET",
			Description: 	"See A Block",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
		case "POST":
			var addBlockBody addBlockBody
			err := json.NewDecoder(r.Body).Decode(&addBlockBody)
			utils.HandleErr(err)
			blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
			rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	// 1. get map
	vars := mux.Vars(r)
	
	// 2. get path parmeter
	height := vars["height"]

	// 3. conversion
	_height, err := strconv.Atoi(height)
	utils.HandleErr(err)

	// 4. get block
	block, err := blockchain.GetBlockchain().GetBlock(_height)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		rw.WriteHeader(http.StatusNotFound)
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	// 이게 adapter 패턴이다.
	// 이와 같이 http HandleFunc는 타입이다.
	// 함수가 호출되면 무언가를 반환하거나 실행하는 것 이 아니다. http HandleFunc 라는 타입이 반환되는 것 이다.

	// Handler는 interface이다.

	// 여기서 추가적으로 http.Handler가 요구하는 serveHTTP가 구현된 객체가 반환된다는 것 이다.
	// 이게 adapter 패턴이다. 
	/**
	func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
		f(w, r)
	}
	**/
	// 추가적으로 보면 serveHTTP가 receiver로 구현되어 있다.
	// 우리가 URL에 marshal text를 사용하기 위해서 URL type을 구성하고 String receiver를 만든 것 처럼
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	router := mux.NewRouter()

	port = fmt.Sprintf(":%d", aPort)

	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")

	fmt.Printf("Listening on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}