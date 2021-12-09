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
		initial("https://bsc-dataseed.binance.org")
	// createWalletKeyStore("xxxx")
	// importKs("0x02178dA504cd690ECD8a444348cF80F711c2a093", "xxxx")
	// updateKs("0x02178dA504cd690ECD8a444348cF80F711c2a093", "xxxx", "yyyy")
}
