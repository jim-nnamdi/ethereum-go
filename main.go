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
	fmt.Printf("balance: %d\n", getBalance)

	txByHash, pending, err := client.TransactionByHash(context.Background(), common.HexToHash("0x74f6e6c8bef01a893c2f9d955513c736db999b07c6461bf3bc221e07ca8b511a"))
	if err != nil {
		log.Printf("error fetching transaction: %v\n", err)
		return
	}
	if !pending {
		if txByHash != nil {
			fmt.Printf("txByHash chainID: %v\n", txByHash.ChainId())
			fmt.Printf("txByHash Nonce: %v\n", txByHash.Nonce())
		}
	}

	transactionReceipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x74f6e6c8bef01a893c2f9d955513c736db999b07c6461bf3bc221e07ca8b511a"))
	if err != nil {
		log.Printf("receipt error: %v", err)
		return
	}
	if transactionReceipt != nil {
		fmt.Printf("tx status:%v\n", transactionReceipt.Status)
		fmt.Printf("contract addr: %v\n", transactionReceipt.ContractAddress)
		fmt.Printf("txByHash blockno: %v\n", transactionReceipt.BlockNumber)
	}

	blockNumber, err := client.BlockByNumber(context.Background(), transactionReceipt.BlockNumber)
	if err != nil {
		log.Printf("cannot fetch block: %v\n", err)
		return
	}
	for _, tx := range blockNumber.Transactions() {
		if tx != nil {
			fmt.Println("gas", tx.Gas())
			fmt.Println("hash", tx.Hash())
			fmt.Println("recipient", tx.To())
		}
	}
}
