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

func block(rw http.ResponseWriter, r *http.Request) {
	// 1. get map
	vars := mux.Vars(r)
	
	// 2. get path parmeter
	height := vars["height"]

	// 3. conversion
	_height, err := strconv.Atoi(height)
	utils.HandleErr(err)

	// 4. get block
	block := blockchain.GetBlockchain().GetBlock(_height)

	json.NewEncoder(rw).Encode(block)
}

func Start(aPort int) {
	router := mux.NewRouter()

	port = fmt.Sprintf(":%d", aPort)

	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")

	fmt.Printf("Listening on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}