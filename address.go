//go:build ignore

package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// make an address from a hex string
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	// print the address in different formats
	fmt.Println("address:", address.Hex())                                  // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F print the address in hex format length
	fmt.Println("address hash:", common.BytesToHash(address.Bytes()).Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f print the address in hash format length 66
	fmt.Println("address bytes:", address.Bytes())                          // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111] print the address in bytes format
}
