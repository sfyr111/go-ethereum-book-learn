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
func createKeystore() string {
	// Step 1: Initialize a keystore at the "./tmp" directory
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

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

	return account.URL.Path
}

// Function to import an existing keystore file
func importKeystore() error {
	// 1. specify the source file path
	sourceFile := "./tmp/UTC--2025-02-07T10-11-37.111837000Z--56475063661ed425d87ea1b7b2f41b82ea6c72bd"

	// 2. use different directory to store imported account
	ks := keystore.NewKeyStore("./imported_wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	// 3. read the keystore file content
	jsonBytes, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	// Step 4: Import the account using the keystore JSON content and password
	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Imported Account Address:", account.Address.Hex())

	// Step 5: Delete the old keystore file after import to avoid duplicates
	if err := os.Remove(sourceFile); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("success: %s\n", account.Address.Hex())
	fmt.Printf("source file deleted: %s\n", sourceFile)
	return nil
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
	createKeystore()

	// Run the keystore import function
	importKeystore()

	// Example usage
	// keystorePath := "./tmp/UTC--2024-11-04T10-03-36.699031000Z--57ff2c32a58dfddb4d58cad6df13025c74249ca0"
	// password := "secret"
	// loadAccountFromKeystore(keystorePath, password)
}
