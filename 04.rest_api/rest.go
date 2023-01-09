package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL 		string `json:"url"`
	Method 		string `json:"method"`
	Description string `json:"description"`
	Payload		string `json:"payload,omitempty"`
	AdminMsg	string `json:"-"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL: 			"/",
			Method: 		"GET",
			Description: 	"See Documentation",
			AdminMsg: 		"This is hide field",
		},
		{
			URL: 			"/blocks",
			Method: 		"POST",
			Description: 	"Add A Block",
			Payload: 		"data:string",
			AdminMsg: 		"This is hide field",
		},
	}
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

func Open() {
	http.HandleFunc("/", documentation)

	fmt.Printf("Listening on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
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