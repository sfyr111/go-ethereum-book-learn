//go:build ignore

package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// Step 1: Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Step 2: Convert the private key to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)

	// Step 3: Convert the private key to hexadecimal and print (without "0x" prefix)
	fmt.Println("Private Key (hex):", hexutil.Encode(privateKeyBytes)[2:])

	// Step 4: Get the public key from the private key
	publicKey := privateKey.Public()

	// Step 5: Cast the public key to ECDSA format
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	// Step 6: Convert the public key to bytes
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// Step 7: Convert the public key to hexadecimal and print (without "0x" and EC prefix "04")
	fmt.Println("Public Key (hex):", hexutil.Encode(publicKeyBytes)[4:])

	// Step 8: Generate the Ethereum address from the public key and print
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Ethereum Address:", address)

	// Step 9: Manually hash the public key using Keccak256 and print the Ethereum address
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:]) // Exclude the EC prefix "04"
	fmt.Println("Manual Ethereum Address (hex):", hexutil.Encode(hash.Sum(nil)[12:]))
}
