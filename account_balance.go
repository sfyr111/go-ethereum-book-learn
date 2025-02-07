package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to the Ethereum network using Cloudflare's public gateway
	// client, err := ethclient.Dial("https://cloudflare-eth.com")
	client, err := ethclient.Dial("https://mainnet.base.org")
	if err != nil {
		log.Fatal(err)
	}

	// Specify the Ethereum account address
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	// Fetch the account balance at the latest block (in Wei)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest balance (Wei):", balance) // Output balance in Wei

	// Fetch the account balance at a specific block
	blockNumber := big.NewInt(5532993) // Block number to query balance at
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance at block 5532993 (Wei):", balanceAt) // Output balance at block 5532993

	// Convert balance from Wei to Ether (ETH)
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())                                 // Convert the balance to big.Float for division
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18))) // Convert Wei to ETH
	fmt.Println("Balance at block 5532993 (ETH):", ethValue)               // Output balance in ETH

	// Fetch the pending balance (balance in pending state)
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pending balance (Wei):", pendingBalance) // Output pending balance in Wei
}
