package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	abi "go-eth/abi_build"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func (crypto *Crypto) accountBalance(token_address string) {
	balance, err := crypto.client.BalanceAt(context.Background(), common.HexToAddress(token_address), nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)
}

func (crypto *Crypto) balanceOf(wallet_address string, token_address string) {
	tokenAddress := common.HexToAddress(token_address)
	instance, err := abi.NewToken(tokenAddress, crypto.client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(wallet_address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)
	fmt.Printf("symbol: %s\n", symbol)
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
