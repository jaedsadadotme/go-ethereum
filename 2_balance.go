package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (crypto *Crypto) account(token_address string) {
	address := common.HexToAddress(token_address)

	fmt.Println(address.Hex())                // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(address.Hash().Hex())         // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes())              // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
	fmt.Println("==========================") // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
	crypto.address = address
}

func (crypto *Crypto) accountBalance() {
	fmt.Println(crypto.address)
	balance, err := crypto.client.BalanceAt(context.Background(), crypto.address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue) // 25.729324269165216041
}
