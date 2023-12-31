package solana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
  // "github.com/mr-tron/base58"
)

type RequestBody struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int64       `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
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

type RequestAirdropResponse struct {
  Jsonrpc string `json:"jsonrpc"`
  Result string `json:"result"`
  ID uint64 `json:"id"` 
}

// global variable
var rpc_cluster = ""

const (
  devnet = "https://api.devnet.solana.com" 
  testnet = "https://api.mainnet-beta.solana.com" 
)

func isConnectedToValidRpc() error {
  if rpc_cluster == devnet || rpc_cluster == testnet {
    return nil
  }

  return errors.New("initiate rpc cluster before making request")
}

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
		return errors.New("invalid argument provided")
	}
}

func GetBalance(s string) (uint64, error) {
  if err := isConnectedToValidRpc(); err != nil {
    log.Fatalf("Failed to make request, initiate a rpc_cluster first")
  }

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
	responseBody, _:= io.ReadAll(resp.Body)
	if err := json.Unmarshal(responseBody, &r); err != nil {
		return 0, err
	}

	return r.Result.Value, nil
}

type ParamsOpt struct {
	Commitment string `json:"commitment"`
}

func GetAccountInfo() uint64 {
  if err := isConnectedToValidRpc(); err != nil {
    log.Fatalf("Failed to make request, initiate a rpc_cluster first")
  }

	return 0
}

func GetLatestBlockHash() GetLatestBlockHashValue {
  if err := isConnectedToValidRpc(); err != nil {
    log.Fatalf("Failed to make request, initiate a rpc_cluster first")
  }

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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("asdasdad")
	}

	var r GetLatestBlockHashResponse
	err = json.Unmarshal(data, &r)
	if err != nil {
		fmt.Println("Cannot unmarshal")
	}

	return r.Result.Value
}

  // {
  //   "jsonrpc": "2.0", "id": 1,
  //   "method": "requestAirdrop",
  //   "params": [
  //     "83astBRguLMdt2h5U1Tpdq5tjFoJ6noeGwaY3mDLVcri",
  //     1000000000
  //   ]
  // }

type RequestAirdropParams struct {
  Jsonrpc string `json:"jsonrpc"`
  ID uint64 `json:"id"`
  Method string `json:"method"`
  Params []interface{} `json:"params"`
}

func RequestAirdrop(id uint64, pubkey string, lamport uint64) (RequestAirdropResponse, error) {
  if err := isConnectedToValidRpc(); err != nil {
    log.Fatalf("Failed to make request, initiate a rpc_cluster first")
  }
	
  var params []interface{}
  params = append(params, pubkey, lamport)
  args := RequestAirdropParams{"2.0", id, "requestAirdrop", params}
  
  requestBody, err := json.Marshal(args)
  if err != nil {
    log.Fatalf("Error when trying to marshal req body in request airdrop")
  }

  resp, err := http.Post(rpc_cluster, "application/json", bytes.NewBuffer(requestBody))
  if err != nil {
    log.Fatalf("Somethign went wrong in the post request: %v", err)
  }

  data, err:= io.ReadAll(resp.Body)
  if err != nil {
    log.Fatalf("Something went wrong, error: %v", err)
  }

  var response RequestAirdropResponse 
  err = json.Unmarshal(data, &response)
  if err != nil {
    log.Fatalf("Something went wrong, error: %v", err)
  }

  stringReader := strings.NewReader("Some string, this method turn string into bytes")
  
  fmt.Printf("String reader: %v %T", stringReader, stringReader)
  b, _ := io.ReadAll(stringReader)

  integer, _ := stringReader.Read(b)

  fmt.Printf("b: %v %s %T", b,b,b)
  fmt.Printf("integer: %v", integer)


  return response, nil
}

type PublicKey [32]byte

type Instruction struct {
  ProgramID uint16
  Accounts []uint16
  InstructionData []byte
}

type Address [32]byte

type Message struct {
  Header []byte // contains three unsigned 8-bit values
  
  Addresses Address   
  // read-write accounts first
  // read-only accounts following

  RecentBlockhash [32]byte
  Instructions []Instruction
}

type Signature [64]byte

type Transaction struct {
  Signatures []Signature 
  Message Message
}

// How to serialize instruction data on the client

// Build instruction payload
// Serialize instruction payload

