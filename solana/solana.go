// solana.GetBalance
// solana.Connect("devnet")
// solana.Connect("mainnet")
package solana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Body struct {
	Jsonrpc string   `json:"jsonrpc"`
	ID      int64    `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type ResponseBody struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
	ID      int64  `json:"id"`
}

type Result struct {
	Context interface{} `json:"context"`
	Value   int64       `json:"value"`
}

// global variable
var rpc_cluster = ""

// err := solana.Connect("")
// if err != nil, do something
func Connect(s string) error {
	myPointer := &rpc_cluster
	switch s {
	case "devnet":
		*myPointer = "https://api.devnet.solana.com"
		return nil
	case "mainnet":
		*myPointer = "https://api.mainnet-beta.solana.com"
		return nil
	default:
		return errors.New("Invalid argument provided")
	}
}

// solana.GetBalance("5CXH8Kqhh6f9Gee6GUfsc7VVCbDSSN2NU2x1WyEdNyic")
func GetBalance(s string) (ResponseBody, error) {
	var a = Body{"2.0", 1, "getBalance", []string{s}}
	jsonBody, err := json.Marshal(&a)
	if err != nil {
		fmt.Println("a ...any")
		return ResponseBody{}, err
	}

	resp, err := http.Post(rpc_cluster, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("error")
		return ResponseBody{}, err
	}

	var r ResponseBody
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(responseBody, &r); err != nil {
		fmt.Println("erorr")
		return ResponseBody{}, err
	}

	return r, nil
}
