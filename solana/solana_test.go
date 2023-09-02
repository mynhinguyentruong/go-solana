package solana

import (
  "testing"
  "fmt"
)

func TestRequestAirdrop(t *testing.T) {
  if err := Connect("devnet"); err != nil {
    t.Errorf("Error when connecting: %v", err)
	}
	fmt.Println("Successfully connect")

  resp, err := RequestAirdrop(2423532543, "CXCmMYJdfYYhqjv3K8vUBqx9GvSJHJyv5KWegu9JUzvT", 1000000)

  if err != nil {
    t.Errorf("Something went wrong")
  }

  fmt.Printf("Response: %v", resp)
} 
