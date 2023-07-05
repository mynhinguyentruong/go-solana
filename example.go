package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/mynhinguyentruong/go-solana/solana"
)

type Server struct {
	Port string
}

type Solana struct {
	Rpc_url string
}

type Rpc_Url struct {
	Devnet string
}

type RequestBody struct {
	Jsonrpc string   `json:"jsonrpc"`
	ID      int      `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}
type Result struct {
	Value int64 `json:"value"`
}

type ResponseBody struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Result `json:"result"`
}

// solana.GetBalance("5CXH8Kqhh6f9Gee6GUfsc7VVCbDSSN2NU2x1WyEdNyic")

func main() {
	if err := solana.Connect("devnet"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connect")
	returnValue, err := solana.GetBalance("5CXH8Kqhh6f9Gee6GUfsc7VVCbDSSN2NU2x1WyEdNyic")
	fmt.Println("asdasdasd: ", returnValue.Result.Value)
	fmt.Println("asdasdasd: ", returnValue.Result.Value)
	fmt.Println("asdasdasd: ", returnValue.Result.Value)
	myChannel := make(chan int)

	var wg sync.WaitGroup
	str := []string{"5CXH8Kqhh6f9Gee6GUfsc7VVCbDSSN2NU2x1WyEdNyic"}

	reqBody := RequestBody{"2.0", 1, "getBalance", str}

	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://api.mainnet-beta.solana.com", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	// call to wait for all data coming from resp
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	var b ResponseBody

	if err := json.Unmarshal(responseBody, &b); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response body: ", b.Result.Value)

	a := <-myChannel

	fmt.Println(a)

	wg.Wait()
}

func sendMessage(c chan<- int, i int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %v started working", i)
	c <- i
	defer wg.Done()
}

type Cat struct {
	Name string
}

// Implement GetName method on Cat type
func (cat *Cat) GetName() string {
	return cat.Name
}
