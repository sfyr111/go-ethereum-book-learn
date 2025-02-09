//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Step 1: Connect to the Ethereum client
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client:", err)
	}

	// Step 2: Get the latest block header
	header, err := client.HeaderByNumber(context.Background(), nil) // `nil` for the latest block
	if err != nil {
		log.Fatal("Failed to retrieve block header:", err)
	}
	fmt.Println("Latest Block Number:", header.Number.String())

	// Step 3: Get a specific block by its number (5671744 in this example)
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("Failed to retrieve block:", err)
	}

	// Step 4: Display block details
	fmt.Println("Block Number:", block.Number().Uint64())        // Block number
	fmt.Println("Timestamp:", block.Time())                      // Block timestamp (seconds since epoch)
	fmt.Println("Difficulty:", block.Difficulty().Uint64())      // Block difficulty
	fmt.Println("Block Hash:", block.Hash().Hex())               // Block hash (unique identifier)
	fmt.Println("Transaction Count:", len(block.Transactions())) // Number of transactions in the block

	// Step 5: Get the transaction count using the block's hash
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal("Failed to retrieve transaction count:", err)
	}
	fmt.Println("Transaction Count (using TransactionCount):", count)
}
