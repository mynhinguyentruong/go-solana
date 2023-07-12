package solana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestBody struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int64       `json:"id"`
	Method  string      `json:"method"`
	Params  []ParamsOpt `json:"params"`
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

type GetLatestBlockHashResponse struct {
	Jsonrpc string                   `json:"jsonrpc"`
	Result  GetLatestBlockHashResult `json:"result"`
	ID      uint64                   `json:"id"`
}

type GetLatestBlockHashResult struct {
	Context interface{}             `json:"context"`
	Value   GetLatestBlockHashValue `json:"value"`
}

type GetLatestBlockHashValue struct {
	Blockhash            string `json:"blockhash"`
	LastValidBlockHeight uint64 `json:"lastValidBlockHeight"`
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

func GetBalance(s string) (uint64, error) {
	var a = RequestBody{"2.0", 1, "getBalance", []ParamsOpt{
		{"processed"},
	}}
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

type ParamsOpt struct {
	Commitment string `json:"commitment"`
}

func GetAccountInfo() uint64 {
	return 0
}

func GetLatestBlockHash() GetLatestBlockHashValue {
	a := RequestBody{"2.0", 1, "getLatestBlockhash", []ParamsOpt{
		{"processed"},
	}}

	jsonBody, err := json.Marshal(a)
	if err != nil {
		log.Fatal("v ...any")
	}

	resp, err := http.Post(rpc_cluster, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal("sdasd")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("asdasdad")
	}

	var r GetLatestBlockHashResponse
	err = json.Unmarshal(data, &r)
	if err != nil {
		fmt.Println("Cannot unmarshal")
	}

	fmt.Printf("Data: %v", r)

	return r.Result.Value

}
