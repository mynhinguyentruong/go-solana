package main

import (
	"fmt"
	"log"

	"github.com/mynhinguyentruong/go-solana/solana"
)

func main() {
	if err := solana.Connect("devnet"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connect")
	returnValue, err := solana.GetBalance("5CXH8Kqhh6f9Gee6GUfsc7VVCbDSSN2NU2x1WyEdNyic")
	if err != nil {
		log.Fatal(err)
	}
	balance := float64(returnValue) / 1000000000

	fmt.Printf("The current balance is: %f SOL", balance)

}
