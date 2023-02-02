package custom

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func GetBalanceAtBlock(blockNumber string) (*big.Int, error) {
	conn, err := ChainConnect()
	if err != nil {
		log.Println("connection error", err)
	}
	balance, err := conn.BalanceAt(context.Background(), common.HexToAddress(blockNumber), nil)
	if err != nil {
		log.Printf("balance error: %v\n", err)
		return nil, err
	}
	return balance, nil
}
