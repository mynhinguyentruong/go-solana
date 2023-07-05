// solana.GetBalance
// solana.Connect("devnet")
// solana.Connect("mainnet")
package solana

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type RequestBody struct {
	Jsonrpc string   `json:"jsonrpc"`
	ID      int64    `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type ResponseBody struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
	ID      uint64 `json:"id"`
}

type Result struct {
	Context interface{} `json:"context"`
	Value   uint64      `json:"value"`
}

// global variable
var rpc_cluster = ""

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
func GetBalance(s string) (uint64, error) {
	var a = RequestBody{"2.0", 1, "getBalance", []string{s}}
	jsonBody, err := json.Marshal(&a)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(rpc_cluster, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, err
	}

	var r ResponseBody
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(responseBody, &r); err != nil {
		return 0, err
	}

	return r.Result.Value, nil
}
