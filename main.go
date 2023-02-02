// Copyright 2020 Jim Nnamdi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jim-nnamdi/ethereum-go/custom"
)

func main() {
	var (
		transactionHashString = "0x74f6e6c8bef01a893c2f9d955513c736db999b07c6461bf3bc221e07ca8b511a"
	)

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/5c0531336045410e9bb1c1e0d3fec3eb")
	if err != nil {
		log.Printf("cannot connect: %v", err)
		return
	}

	// get transaction receipts
	transactionReceipt, err := custom.GetTransactionReceipt(transactionHashString, client)
	if err != nil || transactionReceipt == nil {
		log.Printf("error fetching tx receipt: %v\n", err)
		return
	}
	fmt.Printf("tx status:%v\n", transactionReceipt.Status)
	fmt.Printf("contract addr: %v\n", transactionReceipt.ContractAddress)
	fmt.Printf("txByHash blockno: %v\n", transactionReceipt.BlockNumber)

	// fetch balance
	balance, err := custom.GetBalanceAtBlock("0xbd3Afb0bB76683eCb4225F9DBc91f998713C3b01", client)
	if err != nil {
		log.Printf("balance error: %v\n", balance)
		return
	}
	fmt.Println("balance:", balance)

	// fetch transactions from hash.
	custom.GetTransactions(context.Background(), transactionHashString, client)
}
