package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/5c0531336045410e9bb1c1e0d3fec3eb")
	if err != nil {
		log.Printf("cannot connect: %v", err)
		return
	}
	getBalance, err := client.BalanceAt(context.Background(), common.HexToAddress("0xbd3Afb0bB76683eCb4225F9DBc91f998713C3b01"), nil)
	if err != nil {
		log.Printf("error fetching balances: %v", err)
		return
	}
	fmt.Printf("balance: %d", getBalance)
}
