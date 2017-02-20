# utxohunt

### Office hours for this lab: Tuesday, Feb 21st, 3 to 6pm, E15-357 (Media lab DCI office)
(Note -- I didn't realize Tuesday was on a Monday schedule; moved office hours a bit later)

For the first assignment, we're going to make some transactions on the Bitcoin test network.  We're using btcd as our libraries, which is bitcoin written in golang.  The goal is to understand how transactions are constructed and signed, and to become familiar with the utxo model bitcoin uses.

Testnet3 is a network for testing out bitcoin.  It works almost exactly like the regular bitcoin network (small changes to addresses, the difficulty of proof of work) but everyone agrees that the testnet coins are not worth anything.  This isn't enforced by anything on the network, it's just something people decide.  The fact that it's testnet3 indicates that this rule failed for testnets 1 and 2, when people started trading the testnet coins for mainnet coins.

As we went over in class, the bitcoin ledger is built out of transactions; transactions are really all there is.  Transactions are grouped into blocks, and blocks are linked together in a chain, giving a universally agreed upon sequence of transactions.  Transactions have two components: inputs and outputs.  The coins exist as unspent transaction outputs, or UTXOs.  When outputs are spent (as inputs) they are completely consumed.

In this lab you'll be performing many of the functions of wallet software, by identifying outputs to spend, creating transactions, signing them, and broadcasting them to the network.  Most wallet software does this all automatically, but this assignment is more manual so you can see how it works.


## Setup

[Get go installed](https://golang.org/), either on Athena or your local linux machine.  Mac and Windows will probably work, but are not supported (Don't ask us how to get things working on mac/win)

Once go is running and you have a $GOPATH set, get the class repo:

$ go get -v github.com/mit-dci/utxohunt

In this repo there are 4 files:
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

First, look in utxohunt/addrfrompriv.go, and make a keypair.  The AddressFromPrivateKey() function will help you.  Put your own random string in to generate a private key.  If you call the AddressFromPrivateKey() function, it will return that address as a string, as well as give you the compressed public key and pay to witness pubkey hash script.

Save this address (it starts with an "m" or "n").  You'll need this to send the money to yourself.

## Task 2: Find the first treasure hunt transaction

A _block explorer_ is a website which watches the blockchain and parses out information about blocks, addresses, and transactions.  You can use this blockexplorer to see what's happening on the Bitcoin testnet: [https://testnet.smartbit.com.au/](https://testnet.smartbit.com.au/).

We created a transaction with one hundred outputs.
`b3975fe93f93d028bcc5fb1c1a3f7c1b77c9a558ace98edeba27be6904fcc61b` is the "txid" or unique identifier of the transaction.  (The txid is the hash of the serialized transaction)

This transaction has many outputs.  The outputs are on the right side, and the inputs (witness_v0_keyhash) are on the left side.

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


## Further steps / bonus money

Try to get some more money.  This is doable with just the block explorer.  There are more unspent outputs on the network to find.  One has a private key which is the double-sha256 of the *address* from which you took the first coins.  

To grab these coins, you will need to use AddressFromPrivateKey() to generate that address, search the blockchain for the txid, and try to send an output to yourself, the same way as with the first transaction you created.

There are 2 more outputs, both secured with private keys that are double-sha256s of 2-digit numbers (ascii-encoded).  Try to find them and grab them.

Harder to find:  The two-digit private key outputs form a chain, with the change of the first transaction leading to the next transaction.  After these two outputs were created, 3.35968754 coins (335968754 satoshis) were sent back to a witness_v0 address.

The txid of this final transaction with 1 input and 1 output is the private key for even more coins.
Note that in bitcoin, txids are displayed backwards!  For whatever reason, the actual sha256 output has all its bytes reversed before being shown or saved to disk.  Functions for dealing with reversed txids are in the btcd libraries, so this should be possible.

Note that in many cases, someone else in the class may have grabbed the coins before you.  That's OK, just write down where you found the coins to be and the private key you would have used to take them.


## Submitting work

Submit your homework... ON THE BLOCKCHAIN!
Use OP_RETURN!

The opreturn.go file will walk you through how to make an OP_RETURN transaction.

These transactions are on the public blockchain, and we'll find them there.  No need for e-mail or file attachments!

OP_RETURN outputs start with a single opcode (OP_RETURN) which terminates script execution.  So the output can never be spend, and is thus not added to the utxo database.  You can put whatever data you want after the OP_RETURN, though it's limited to 40 bytes in length.

For this assignment, sending your coins to an OP_RETURN output with your name or MIT ID number is cryptographic proof that you sent the coins (or someone else did, impersonating you!)

Use opreturn.go to create transactions FROM the transactions you sent to yourself using EZTxBuilder().  The created transaction will send from and to the same address, adding an OP_RETURN output.  Broadcast this to the network, hope it gets into a block, and you're done! 

To parse other OP_RETURNs, or the one you made yourself, try using python.

Here's an example transaction:
`c29dc7b974901989c156578fc8dd341752bf28e415191bb1dc4b3aabc3ac11c5`
the OP_RETURN is 363839322054657374206f7574707574

Load up python in your terminal (most linux and mac terminals have it) by running ` $ python `
from there:
``` >>> "363839322054657374206f7574707574".decode("hex")
'6892 Test output' ```

Prefix all your OP_RETURNs with 6892 so it's easy to search for them.

If you only grab a little bit of money and send an OP_RETURN, that's fine.  If you manage to get some of the bonus utxos and send OP_RETURNs, even better!  If you want to get really fancy, try aggregating all your outputs into a single, higher value tx output. (Code left as excercise to the reader)
