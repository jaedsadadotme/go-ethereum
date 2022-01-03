package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	abi "go-eth/abi_build/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (crypto *Crypto) account() {
	address := crypto.address

	fmt.Println(address.Hex())                // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(address.Hash().Hex())         // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes())              // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
	fmt.Println("==========================") // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
	crypto.address = address
}

func (crypto *Crypto) accountBalance(token_address string) {
	balance, err := crypto.client.BalanceAt(context.Background(), crypto.address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)
}

func (crypto *Crypto) balanceOf() {
	for _, token := range tokens_list {
		tokenAddress := common.HexToAddress(token)
		instance, err := abi.NewToken(tokenAddress, crypto.client)
		if err != nil {
			log.Fatal(err)
		}

		address := crypto.address
		bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
		if err != nil {
			log.Fatal(err)
		}

		name, err := instance.Name(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		decimals, err := instance.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		fbal := new(big.Float)
		fbal.SetString(bal.String())
		value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

		fmt.Printf("%s balance: %f \n", name, value) // "balance: 74605500.647409"
	}
}
