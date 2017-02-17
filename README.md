# bitgorand

For the first assignment, we're going to make some transactions on the Bitcoin test network.  We're using btcd as our libraries, which is bitcoin written in golang.  The goal is to understand how transactions are constructed and signed, and to become familiar with the utxo model bitcoin uses.

Testnet3 is a network for testing out bitcoin.  It works almost exactly like the regular bitcoin network (small changes to addresses, the difficulty of proof of work) but everyone agrees that the testnet coins are not worth anything.  This isn't enforced by anything on the network, it's just something people decide.  The fact that it's testnet3 indicates that this rule failed for testnets 1 and 2, when people started trading the testnet coins for mainnet coins.

As we went over in class, the bitcoin ledger is built out of transactions; transactions are really all there is.  Transactions are grouped into blocks, and blocks are linked together in a chain, giving a universally agreed upon sequence of transactions.  Transactions have two components: inputs and outputs.  The coins exist as unspent transaction outputs, or UTXOs.  When outputs are spent (as inputs) they are completely consumed.

In this lab you'll be performing many of the functions of wallet software, by identifying outputs to spend


## Setup

[Get go installed](https://golang.org/), either on Athena or your local linux machine.  Mac and Windows will probably work, but are not supported (Don't ask us how to get things working on mac/win)

TODO: Figure out submission.  It's easier if you have a github account if you don't already have one.  If you don't have and don't want a github account submission to a local git server is also OK.

Once go is running and you have a $GOPATH set, get the class repo:

$ go get -v github.com/mit-dci/6.892-public

In this repo, there is a folder called `hw`, in which there's `utxohunt`, which contains 4 files:
```
main.go
addrfrompriv.go  
eztxbuilder.go
opreturn.go
```

Here's what they do

#### main.go
This is the main function which is called when you run `./utxohunt`

Edit this file to call functions from other files when you run the program.

#### addrfrompriv.go
Creates a public key and bitcoin address from a private key.  Addresses are copy&pastable encodings of public key hashes.

#### eztxbuilder.go
Puts a transaction together, signs it, and prints the tx hex to the screen.  This can then be copy & pasted to a block explorer like [https://testnet.smartbit.com.au/txs/pushtx](https://testnet.smartbit.com.au/txs/pushtx), or to your own bitcoin node with `./bitcoin-cli pushrawtransaction (tx hex)`

#### opreturn.go
Similar to eztxbuilder.go, but creates a transaction with 1 input, and 2 outputs.  1 of the outputs is an "OP_RETURN" output which can contain arbitrary data.  Use this to submit your results to the blockchain.

## Task 1: Create a Bitcoin Address

First, look in utxohunt/main.go, and make a keypair.  The AddressFromPrivateKey() function will help you.  Put your own random string in to generate a private key.  If you call the AddressFromPrivateKey() function, it will return that address as a string, as well as give you the compressed public key and pay to witness pubkey hash script.

Save this address (it starts with an "m").  You'll need this to send the money to yourself.

## Task 2: Find the first treasure hunt transaction

A _block explorer_ is a website which watches the blockchain and parses out information about blocks, addresses, and transactions.  You can use this blockexplorer to see what's happening on the Bitcoin testnet: [https://testnet.smartbit.com.au/](https://testnet.smartbit.com.au/).

We created a transaction with hundreds of outputs.
a56f06aa7dbc3b3ba86390f163f3639e962d3e223e9ca070002288eb555bb830 is
the "txid" or unique identifier of the transaction.  (The txid is the
hash of the serialized transaction)

This transaction has many outputs.  The outputs are on the right side, and the single input (witness_v0_keyhash) is on the left side.

Outpoints are defined by a txid and the output number.  On the block explorer page, the output on the top right is output 0, then 1, 2, etc.

The outputs are all sent to the same address.  The private key which led to this script is the
double-sha256 of "BTC secret key MIT".

You're going to find and claim an unspent output in this transaction.  Please be nice and leave the rest of them for your other classmates!

## Task 3: Make a transaction

Using EZTxBuilder(), make a transaction sending from the up-for-grabs transaction to your own address.

You will need to modify 

	hashStr

	outPoint (output index number)

	sendToAddressString

	prevAddressString (the address of the "BTC secret key" pubkey)

	wire.NewTxOut (change the amount to less than the input amount.  A few thousand less is enough of a fee)


When you modify the code, you need to re-compile the code.  Run "go build" in that directory to compile.

You'll get a long hex string which you can test by pasting the transaction into the web interface [https://testnet.smartbit.com.au/txs/pushtx](https://testnet.smartbit.com.au/txs/pushtx).

If you get an error, it might be for one of the following reasons:

1.  Someone has already claimed the output you are trying to get.  Go back and look at the transaction's page and see if the output is still available.  It will say "inputs spent" or equivalent.	

2.  64: non-mandatory-script-verify-flag (Signature must be zero for
failed CHECK(MULTI)SIG operation).  This means your signature was
invalid.  Often this is because the hash being signed was invalid.
This could be because the previous output you signed and the one you
indicted don't match, the wrong amount is being sent to the
WitnessScript function, or some other invalid data is in the
transaction prior to signing.  An invalid signature can also be caused
by using the wrong key.  In that case, you will usually get this
error:

3.  64: non-mandatory-script-verify-flag (Script failed an OP_EQUALVERIFY operation).  This means you're probably using the wrong key to sign with, as the public key used and public key hash in the previous output script don't match.

4.  TX decode failed.  That means you're missing some characters, or the transaction is otherwise unintelligable to the bitcoin-cli parser.

If everything worked, you'll see a txid returned (64 hex characters; 32 bytes).  That means you got money.  You can use the same EZTxBuilder() to send that money somewhere else.


## Submitting work

Submit your homework... ON THE BLOCKCHAIN!
Use OP_RETURN!

There is a 

To parse other OP_RETURNs, or the one you made yourself, try using python.

Here's an example transaction:
c29dc7b974901989c156578fc8dd341752bf28e415191bb1dc4b3aabc3ac11c5
the OP_RETURN is 363839322054657374206f7574707574

Load up python in your terminal (most linux and mac terminals have it) by running ` $ python `
from there:
``` >>> "363839322054657374206f7574707574".decode("hex")
'6892 Test output' ```

Prefix all your OP_RETURNs with 6892 so it's easy to search for them.
