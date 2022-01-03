package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (crypto *Crypto) initial(network string) *Crypto {
	client, err := ethclient.Dial(network)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	fmt.Println("================================================")
	crypto.client = client
	return crypto
}

func main() {
	client := Crypto{}
	client.
		initial("https://bsc-dataseed.binance.org").
		getPrice()
	// balanceOf("0x000000000000000000000000000000")
	// accountBalance("0x000000000000000000000000000000")
	// createWalletKeyStore("xxxx")
	// importKs("0x000000000000000000000000000000", "xxxx")
	// updateKs("0x000000000000000000000000000000", "xxxx", "yyyy")
}
