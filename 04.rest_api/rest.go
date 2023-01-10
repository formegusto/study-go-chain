package rest_test

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/formegusto/study-go-chain/02.block_chain/blockchain"
	"github.com/formegusto/study-go-chain/04.rest_api/rest"
	"github.com/formegusto/study-go-chain/utils"
)

const port string = ":4000"

type URL string
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s",port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL 		URL    `json:"url"`
	Method 		string `json:"method"`
	Description string `json:"description"`
	Payload		string `json:"payload,omitempty"`
	AdminMsg	string `json:"-"`
}

type AddBlockBody struct {
	Message string
}

// func (u URLDescription) String() string {
// 	return "Hello I'm the URL Description"
// }

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL: 			URL("/"),
			Method: 		"GET",
			Description: 	"See Documentation",
			AdminMsg: 		"This is hide field",
		},
		{
			URL: 			URL("/blocks"),
			Method: 		"POST",
			Description: 	"Add A Block",
			Payload: 		"data:string",
			AdminMsg: 		"This is hide field",
		},
	}
	// fmt.Println(data)
	// 1. hard
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)

	// // fmt.Printf("%s",b)
	// // [{"URL":"/","Method":"GET","Description":"See Documentation"}
	// // 원래는 그냥 문자열이라는 거!
	// rw.Header().Add("Content-Type", "appliation/json")
	// fmt.Fprintf(rw, "%s", b)

	// 2. simple
	rw.Header().Add("Content-Type", "appliation/json")
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
		case "POST":
			var addBlockBody AddBlockBody
			// Pointer를 보내주어야 함
			err := json.NewDecoder(r.Body).Decode(&addBlockBody)
			utils.HandleErr(err)
			// fmt.Println(addBlockBody)
			blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
			rw.WriteHeader(http.StatusCreated)
	}
}

func Open(aPort int) {
	// fmt.Println(URLDescription{
	// 	URL: 			"/",
	// 	Method: 		"GET",
	// 	Description: 	"See Documentation",
	// 	AdminMsg: 		"This is hide field",
	// })
	// // "Hello I'm the URL Description"
	// http.HandleFunc("/", documentation)
	// http.HandleFunc("/blocks", blocks)

	// fmt.Printf("Listening on http://localhost:%s\n", port)
	// log.Fatal(http.ListenAndServe(port, nil))
	rest.Start(aPort)
}

// Documentation Example
/**
* GET
* /
* See Documentation
*/

/**
* POST
* /blocks
* Create a block
*/