package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Crypto struct {
	client  *ethclient.Client
	wallet  string
	address common.Address
}
