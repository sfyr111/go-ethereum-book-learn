package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Step 1: Connect to Ethereum client
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client:", err)
	}

	// Step 2: Retrieve a block by number
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("Failed to retrieve block:", err)
	}

	// Step 3: Iterate over transactions in the block
	for _, tx := range block.Transactions() {
		fmt.Println("Transaction Hash:", tx.Hash().Hex())       // transaction hash
		fmt.Println("Value (Wei):", tx.Value().String())        // transaction value
		fmt.Println("Gas:", tx.Gas())                           // gas
		fmt.Println("Gas Price (Wei):", tx.GasPrice().Uint64()) // gas price
		fmt.Println("Nonce:", tx.Nonce())                       // nonce
		fmt.Println("Data:", tx.Data())                         // data
		fmt.Println("To Address:", tx.To().Hex())               // to address

		// Retrieve chainID for EIP155 signer
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal("Failed to get chain ID:", err)
		}
		// Step 4: Get the sender's address using the transaction's AsMessage method
		signer := types.NewEIP155Signer(chainID)
		if msg, err := types.Sender(signer, tx); err == nil {
			fmt.Println("From Address:", msg.Hex())
		}

		// Step 5: Get the transaction receipt to check the status
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal("Failed to get transaction receipt:", err)
		}
		fmt.Println("Transaction Status:", receipt.Status) // 1 for success, 0 for failure
	}

	// Step 6: Query transactions using block hash and transaction index
	blockHash := block.Hash()
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal("Failed to get transaction count for block:", err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal("Failed to retrieve transaction by index:", err)
		}
		fmt.Println("Transaction Hash (by Index):", tx.Hash().Hex())
	}

	// Step 7: Query a transaction directly by transaction hash
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal("Failed to retrieve transaction by hash:", err)
	}

	fmt.Println("Direct Transaction Hash:", tx.Hash().Hex())
	fmt.Println("Is Transaction Pending:", isPending)
}
