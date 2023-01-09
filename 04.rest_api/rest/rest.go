package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/formegusto/study-go-chain/02.block_chain/blockchain"
	"github.com/formegusto/study-go-chain/utils"
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
			URL: 			url("/blocks/{}"),
			Method: 		"GET",
			Description: 	"See A Block",
		},
	}
	rw.Header().Add("Content-Type", "appliation/json")
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
		case "POST":
			var addBlockBody addBlockBody
			err := json.NewDecoder(r.Body).Decode(&addBlockBody)
			utils.HandleErr(err)
			blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
			rw.WriteHeader(http.StatusCreated)
	}
}

func Start(aPort int) {
	handler := http.NewServeMux()

	port = fmt.Sprintf(":%d", aPort)

	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)

	fmt.Printf("Listening on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}