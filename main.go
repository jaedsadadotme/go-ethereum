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
	crypto.client = client
	return crypto
}

func main() {
	client := Crypto{}
	client.
		initial("https://bsc-dataseed.binance.org").
		balanceOf("0x999999999999999999999999999999999", "0x2170ed0880ac9a755fd29b2688956bd959f933f8")
	// accountBalance("0x999999999999999999999999999999999")
	// createWalletKeyStore("xxxx")
	// importKs("0x999999999999999999999999999999999", "xxxx")
	// updateKs("0x999999999999999999999999999999999", "xxxx", "yyyy")
}
