package solana

import (
  "testing"
  // "crypto/ed25519"
  "fmt"
)

func TestCreateNewKeypair(t *testing.T) {
  keypair := Keypair{}
  keypair.Init()

  if keypair.PublicKey == nil {
    t.Errorf("valid signature rejected")
  }
  
  fmt.Println("Keypair: ", keypair.PublicKey)

}



