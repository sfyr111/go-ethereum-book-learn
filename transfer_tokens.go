package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func main() {
	// Step 1: Connect to Ethereum client
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client:", err)
	}

	// Step 2: Load private key
	privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
	if err != nil {
		log.Fatal("Failed to load private key:", err)
	}

	// Derive sender address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Failed to retrieve account nonce:", err)
	}

	// Step 3: Set transaction parameters
	value := big.NewInt(0) // in wei (0 ETH)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to get suggested gas price:", err)
	}
	toAddress := common.HexToAddress("0xE280029a7867BA5C9154434886c241775ea87e53")    // Recipient address
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d") // ERC-20 contract address

	// Step 4: Calculate method ID for transfer function
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println("Method ID:", hexutil.Encode(methodID)) // 0xa9059cbb

	// Step 5: Prepare data payload
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println("Padded recipient address:", hexutil.Encode(paddedAddress))
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens, assuming 18 decimals
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println("Padded amount:", hexutil.Encode(paddedAmount))

	// Concatenate methodID, recipient address, and amount
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// Step 6: Estimate gas limit
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal("Failed to estimate gas limit:", err)
	}
	fmt.Println("Estimated gas limit:", gasLimit)

	// Step 7: Create transaction
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	// Step 8: Sign transaction
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("Failed to get network ID:", err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("Failed to sign transaction:", err)
	}

	// Step 9: Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Failed to send transaction:", err)
	}

	fmt.Printf("Transaction sent! Hash: %s\n", signedTx.Hash().Hex())
}
