package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Function to create a new keystore account and save it to disk
func createKeystore() {
	// Step 1: Initialize a keystore at the "./wallets" directory
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	// Step 2: Set a password for encrypting the private key
	password := "secret"

	// Step 3: Generate a new account with the keystore, encrypted with the password
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	// Step 4: Print the newly generated Ethereum address
	fmt.Println("New Account Address:", account.Address.Hex())
	fmt.Println("Keystore file:", account.URL.Path)
}

// Function to import an existing keystore file
func importKeystore() {
	// Step 1: Specify the path to the existing keystore JSON file
	// file := "./wallets/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
	file := "./wallets/UTC--2024-11-04T10-01-31.017923000Z--57ff2c32a58dfddb4d58cad6df13025c74249ca0"

	// Step 2: Initialize a new keystore at a different directory
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	// Step 3: Read the JSON content of the keystore file
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// Step 4: Import the account using the keystore JSON content and password
	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	// Step 5: Print the imported account's address
	fmt.Println("Imported Account Address:", account.Address.Hex())

	// Step 6: Delete the old keystore file after import to avoid duplicates
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

// Function to read and display account information
func loadAccountFromKeystore(keystoreFile string, password string) {
	// Step 1: Read the JSON content of the keystore file
	jsonBytes, err := ioutil.ReadFile(keystoreFile)
	if err != nil {
		log.Fatalf("Failed to read keystore file: %v", err)
	}

	// Step 2: Decrypt the keystore using the password
	key, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatalf("Failed to decrypt keystore: %v", err)
	}

	// Step 3: Retrieve the private key as bytes
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private Key (hex):", hex.EncodeToString(privateKeyBytes))

	// Step 4: Retrieve and print the public Ethereum address
	address := key.Address.Hex()
	fmt.Println("Ethereum Address:", address)

	// Optionally, access the public key if needed
	publicKeyECDSA, ok := key.PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}

	// Step 5: the public key in hex format
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public Key (hex):", hexutil.Encode(publicKeyBytes)[4:])
}

func main() {
	// Run the keystore creation function
	// createKeystore()

	// Uncomment the line below to run the keystore import function
	// importKeystore()

	// Example usage
	keystorePath := "./wallets/UTC--2024-11-04T10-03-36.699031000Z--57ff2c32a58dfddb4d58cad6df13025c74249ca0"
	password := "secret"
	loadAccountFromKeystore(keystorePath, password)
}
