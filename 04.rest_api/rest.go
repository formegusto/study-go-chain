package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/formegusto/study-go-chain/utils"
)

const port string = ":4000"

type URLDescription struct {
	URL 		string
	Method 		string
	Description string 
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL: 			"/",
			Method: 		"GET",
			Description: 	"See Documentation",
		},
	}
	b, err := json.Marshal(data)
	utils.HandleErr(err)

	// fmt.Printf("%s",b)
	// [{"URL":"/","Method":"GET","Description":"See Documentation"}
	// 원래는 그냥 문자열이라는 거!
	rw.Header().Add("Content-Type", "appliation/json")
	fmt.Fprintf(rw, "%s", b)
}

func Open() {
	http.HandleFunc("/", documentation)

	fmt.Printf("Listening on http://localhost:%s", port)
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