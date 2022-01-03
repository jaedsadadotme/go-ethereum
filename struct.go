package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var tokens = map[string]string{
	"factory": "0xca143ce32fe78f1f7019d7d551a6402fc5350c73", // pancake_factory
	"ETH":     "0x2170ed0880ac9a755fd29b2688956bd959f933f8", // ETH
	"BUSD":    "0xe9e7cea3dedca5984780bafc599bd69add087d56", // BUSD
	"SHIB":    "0x2859e4544c4bb03966803b044a93563bd2d0dd4d", // SHIB
}

type Crypto struct {
	client  *ethclient.Client
	wallet  string
	address common.Address
}
