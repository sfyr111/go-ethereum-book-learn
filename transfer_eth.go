package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func main() {
	// Step 1: Connect to Ethereum client (Infura, Rinkeby test network)
	client, err := ethclient.Dial("http://127.0.0.1:8545") // ganache local network
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client:", err)
	}

	// Step 2: Load private key (sender's private key in hex format)
	privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
	if err != nil {
		log.Fatal("Failed to load private key:", err)
	}

	// Step 3: Derive public key and sender's address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("Sender Address:", fromAddress.Hex())

	// Step 4: Retrieve the nonce for the sender's account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Failed to retrieve account nonce:", err)
	}

	// Step 5: Define the amount of ETH to transfer (1 ETH in wei)
	value := big.NewInt(1000000000000000000) // 1 ETH in wei

	// Step 6: Set the gas limit and retrieve a suggested gas price
	// gasLimit := uint64(21000) // Standard gas limit for ETH transfer
	gasLimit := params.TxGas // uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to retrieve suggested gas price:", err)
	}

	// Step 7: Define the recipient address
	toAddress := common.HexToAddress("0xE280029a7867BA5C9154434886c241775ea87e53")
	fmt.Println("Recipient Address:", toAddress.Hex())

	// Step 8: Create the transaction (no data field as this is a simple ETH transfer)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// Step 9: Get the chain ID for the Rinkeby test network and sign the transaction
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("Failed to get network ID:", err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("Failed to sign the transaction:", err)
	}

	// Step 10: Send the transaction to the Ethereum network
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Failed to send the transaction:", err)
	}

	fmt.Printf("Transaction sent! Transaction Hash: %s\n", signedTx.Hash().Hex())
}
