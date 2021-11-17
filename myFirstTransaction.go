package main

import (
	"fmt"

	"go-algorand-sdk/crypto"
	"go-algorand-sdk/mnemonic"
)

// TODO: insert additional utility functions here

func main() {
	// Create account
	account := crypto.GenerateAccount()
	passphrase, err := mnemonic.FromPrivateKey(account.PrivateKey)
	myAddress := account.Address.String()

	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
	} else {
		fmt.Printf("My address: %s\n", myAddress)
		fmt.Printf("My passphrase: %s\n", passphrase)
		fmt.Println("--> Copy down your address and passpharse for future use.")
		fmt.Println("--> Once secured, press ENTER key to continue...")
		fmt.Scanln()
	}

	// TODO: insert additional codeblocks here
}
