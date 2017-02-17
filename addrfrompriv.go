package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

func AddressFromPrivateKey() (string, error) {

	// private key is the hash of some string (better to use real randomness
	// or a real KDF but this is OK for class.
	// Put any phrase you want here to make your own private key.
	phraseHash := chainhash.DoubleHashB([]byte("621e444a870decf4cd05e5c18334e35c0a"))

	// make a new private key struct.  Private key structs also have a pubkey in them
	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), phraseHash)

	// print out what it looks like in hex, compressed (x-coordinate, y-sign only)

	fmt.Printf("pubkey is %x\n", priv.PubKey().SerializeCompressed())

	hash160 := btcutil.Hash160(priv.PubKey().SerializeCompressed())

	adr, err := btcutil.NewAddressPubKeyHash(hash160, testnet3Parameters)
	if err != nil {
		return "", err
	}

	script, err := txscript.PayToAddrScript(adr)
	if err != nil {
		return "", err
	}
	fmt.Printf("script is: %x\n", script)
	return adr.String(), nil
}
