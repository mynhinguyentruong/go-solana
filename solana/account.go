package solana

import (
	"crypto/ed25519"
	// "fmt"
	"log"
  "errors"
)

var ErrorMissingKeypair = errors.New("cannot find publickey or privatekey... make sure you initialized them by calling Init() method")

type Keypair struct {
  PublicKey ed25519.PublicKey 
  PrivateKey ed25519.PrivateKey
}

func (key *Keypair) Init() (*Keypair) {
  publickey, privatekey, err := ed25519.GenerateKey(nil)
  if err != nil {
    log.Fatal("error while generate new keypair")
  }

  key.PublicKey = publickey
  key.PrivateKey = privatekey

  return key
}

func (key *Keypair) Sign(message string) ([]byte, error) {
  if key.PublicKey == nil || key.PrivateKey == nil {
    return []byte{}, ErrorMissingKeypair
  }
  msg := []byte(message)

  sig := ed25519.Sign(key.PrivateKey, msg)

  return sig, nil
}

func (key *Keypair) Verify(message string, sig []byte) (bool, error) {
  if key.PublicKey == nil || key.PrivateKey == nil {
    return false, ErrorMissingKeypair
  }
  msg := []byte(message)
  return ed25519.Verify(key.PublicKey, msg, sig), nil
}



