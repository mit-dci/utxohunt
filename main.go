package main

import (
	"bytes"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
)

var (
	testnet3Parameters = &chaincfg.TestNet3Params
)

func main() {
	fmt.Printf("bitgorand class program - student version\n")

	// Task #1 make an address pair
	// Call AddressFrom PrivateKey() to make a keypair

	s, err := AddressFromPrivateKey()
	if err != nil {
		panic(err)
	}
	fmt.Printf("address: %s\n", s)

	// Task #2 make a transaction
	// Call EZTxBuilder to make a transaction
	//	tx := EZTxBuilder()
	//	var buf bytes.Buffer
	//	tx.Serialize(&buf)
	//	fmt.Printf("tx in hex:\n%x\n", buf.Bytes())
	//

	tx := OpReturnTxBuilder()
	var buf bytes.Buffer
	tx.Serialize(&buf)
	fmt.Printf("tx in hex:\n%x\n", buf.Bytes())
	return
}
