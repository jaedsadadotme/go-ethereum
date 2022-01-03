package main

import (
	"fmt"
	abi "go-eth/abi_build/pancake"
	factory_abi "go-eth/abi_build/pancake_factory"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (crypto *Crypto) getPrice(token0 string, token1 string) {
	fac_instance, err := factory_abi.NewToken(common.HexToAddress(tokens["factory"]), crypto.client)
	if err != nil {
		log.Fatal(err)
	}
	//
	pair_addr, err := fac_instance.GetPair(&bind.CallOpts{}, common.HexToAddress(token0), common.HexToAddress(token1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pair Address is:", pair_addr)

	instance, err := abi.NewToken(pair_addr, crypto.client)
	if err != nil {
		log.Fatal(err)
	}
	price, err := instance.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	//
	var price0 float64
	var price1 float64
	if price0, err = strconv.ParseFloat(price.Reserve0.String(), 64); err != nil {
		log.Fatal(err)
	}
	if price1, err = strconv.ParseFloat(price.Reserve1.String(), 64); err != nil {
		log.Fatal(err)
	}
	//
	token_instance, err := abi.NewToken(common.HexToAddress(token1), crypto.client)
	if err != nil {
		log.Fatal(err)
	}
	name, err := token_instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)

	}
	token_price := fmt.Sprintf("%s : $%f", name, (price1 / price0))
	fmt.Println(token_price)
}
