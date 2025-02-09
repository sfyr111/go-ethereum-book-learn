//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Function to validate Ethereum address format using regular expression
func isAddressValid(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$") // Match "0x" followed by exactly 40 hexadecimal characters
	return re.MatchString(address)
}

// Function to check if an Ethereum address is a smart contract or regular user account
func isSmartContract(client *ethclient.Client, address common.Address) bool {
	// Fetch the bytecode at the address. `nil` indicates the latest block.
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve code: %v", err)
	}
	// If bytecode length is greater than 0, it's a smart contract
	return len(bytecode) > 0
}

func main() {
	// Step 1: Validate Ethereum address format
	fmt.Printf("Address 0x323b5d4c32345ced77393b3530b1eed0f346429d is valid: %v\n", isAddressValid("0x323b5d4c32345ced77393b3530b1eed0f346429d"))
	fmt.Printf("Address 0xZYXb5d4c32345ced77393b3530b1eed0f346429d is valid: %v\n", isAddressValid("0xZYXb5d4c32345ced77393b3530b1eed0f346429d"))

	// Step 2: Connect to an Ethereum client (using Cloudflare's free gateway)
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client:", err)
	}

	// Step 3: Check if an address is a smart contract
	// Example 1: Known smart contract address (0x Protocol Token (ZRX))
	contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	fmt.Printf("Address %s is a smart contract: %v\n", contractAddress.Hex(), isSmartContract(client, contractAddress))

	// Example 2: Random user account address
	userAddress := common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	fmt.Printf("Address %s is a smart contract: %v\n", userAddress.Hex(), isSmartContract(client, userAddress))
}
