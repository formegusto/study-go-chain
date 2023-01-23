package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/formegusto/study-go-chain/02.block_chain/blockchain"
	"github.com/formegusto/study-go-chain/07.wallet/wallet"
	"github.com/formegusto/study-go-chain/08.p2p/p2p"
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

type balanceResponse struct {
	Address		string	`json:"address"`
	Balance		int		`json:"balance"`
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type addTxPayload struct {
	To		string
	Amount	int
}

type myWalletResponse struct {
	Address 	string	`json:"address"`
}

type addPeerPayload struct {
	Address,Port	string
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
			URL:			url("/status"),
			Method:			"GET",
			Description: 	"See the Status of the Blockchain",
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
			URL: 			url("/blocks/{hash}"),
			Method: 		"GET",
			Description: 	"See A Block",
		},
		{
			URL: 			url("/balance/{address}"),
			Method: 		"GET",
			Description: 	"Get TxOuts for an address",
		},
		{
			URL: 			url("/ws"),
			Method: 		"GET",
			Description: 	"Upgrade to WebSockets",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			err := json.NewEncoder(rw).Encode(blockchain.Blocks(blockchain.Blockchain()))
			utils.HandleErr(err)
			
		case "POST":
			blockchain.Blockchain().AddBlock()
			rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	hash := vars["hash"]

	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		rw.WriteHeader(http.StatusNotFound)
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		next.ServeHTTP(rw, r)
	})
}

func status(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(blockchain.Blockchain())
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	total := r.URL.Query().Get("total")

	switch total {
		case "true":
			amount := blockchain.BalanceByAddress(address, blockchain.Blockchain())
			json.NewEncoder(rw).Encode(balanceResponse{address, amount})
		default:
			json.NewEncoder(rw).Encode(blockchain.UTxOutsByAddress(address, blockchain.Blockchain()))
	}
}

func mempool(rw http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(rw).Encode(blockchain.Mempool.Txs)
	utils.HandleErr(err)
}

func transactions(rw http.ResponseWriter, r *http.Request) {
	var payload addTxPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	utils.HandleErr(err)

	err = blockchain.Mempool.AddTx(payload.To, payload.Amount)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(errorResponse{err.Error()})
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func myWallet(rw http.ResponseWriter, r *http.Request) {
	address := wallet.Wallet().Address
	json.NewEncoder(rw).Encode(myWalletResponse{address})
	// anonymous struct
	// json.NewEncoder(rw).Encode(struct{ 
	// 	Address	string	`json:"address"` 
	// }{ address })
}

func peers(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			var payload addPeerPayload
			json.NewDecoder(r.Body).Decode(&payload)
			p2p.AddPeer(payload.Address, payload.Port, port)
			rw.WriteHeader(http.StatusOK)
		case "GET":
			json.NewEncoder(rw).Encode(p2p.AllPeers(&p2p.Peers))
	}
}


func Start(aPort int) {
	router := mux.NewRouter()

	port = fmt.Sprintf(":%d", aPort)

	router.Use(jsonContentTypeMiddleware, loggerMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status)
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	router.HandleFunc("/balance/{address}", balance).Methods("GET")
	router.HandleFunc("/mempool", mempool).Methods("GET")
	router.HandleFunc("/wallet", myWallet).Methods("GET")
	router.HandleFunc("/transactions", transactions).Methods("POST")
	router.HandleFunc("/ws", p2p.Upgrade).Methods("GET")
	router.HandleFunc("/peers", peers).Methods("GET","POST")

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}