package main

import (
	"fmt"
	"log"

	"github.com/mynhinguyentruong/go-solana/solana"
)

func main() {

	if err := solana.Connect("mainnet"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connect")
	returnValue, err := solana.GetBalance("5CXH8Kqhh6f9Gee6GUfsc7VVCbDSSN2NU2x1WyEdNyic")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(returnValue)
	balance := float64(returnValue) / 1000000000

	fmt.Printf("The current balance is: %v SOL\n", balance)

	lastestBlockhash := solana.GetLatestBlockHash()

	fmt.Printf("\nBlockhash: %v\nLast Valid Block Heigh: %v\n", lastestBlockhash.Blockhash, lastestBlockhash.LastValidBlockHeight)

  resp, err := solana.RequestAirdrop(24234, "CXCmMYJdfYYhqjv3K8vUBqx9GvSJHJyv5KWegu9JUzvT", 1000000000)
  if err != nil {
    log.Fatalf("error %v", err)
  }
  fmt.Printf("airdrop: %v", resp)
}

type Transaction struct {
	Accounts     []string // specify all account intend to read or write data to
	Metadata     string
	Instructions []Instruction
}

type Instruction struct {
	Keys       []string
	ProgramID  string
	BufferData []byte
}

type Keys struct {
	pubkey     PublicKey
	isSigner   bool
	isWritable bool
}

type PublicKey struct {
	_bn string // hex representation of address
}
