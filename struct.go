package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var tokens = map[string]string{
	"factory": "0xca143ce32fe78f1f7019d7d551a6402fc5350c73", // pancake_factory
	"BTC":     "0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c",
	"ETH":     "0x2170ed0880ac9a755fd29b2688956bd959f933f8", // ETH
	"BUSD":    "0xe9e7cea3dedca5984780bafc599bd69add087d56", // BUSD
	"SHIB":    "0x2859e4544c4bb03966803b044a93563bd2d0dd4d", // SHIB
	"BSC_USD": "0x55d398326f99059ff775485246999027b3197955",
	"ADA":     "0x3ee2200efb3400fabb9aacf31297cbdd1d435d47",
}

var tokens_list = []string{
	"0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c", //BTCB
	"0x2170ed0880ac9a755fd29b2688956bd959f933f8", // ETH
	"0xe9e7cea3dedca5984780bafc599bd69add087d56", // BUSD
	"0x2859e4544c4bb03966803b044a93563bd2d0dd4d", // SHIB
}

type Crypto struct {
	client  *ethclient.Client
	wallet  string
	address common.Address
}
