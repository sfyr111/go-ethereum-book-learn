package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// client, err := ethclient.Dial("wss://mainnet.base.org")
	// client, err := ethclient.Dial("wss://cloudflare-eth.com")
	client, err := ethclient.Dial("wss://eth-mainnet.g.alchemy.com/v2/YOUR-API-KEY")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
