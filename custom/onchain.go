package custom

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// pass in the blockNumber as args
func GetBalanceAtBlock(blockNumber string, conn *ethclient.Client) (*big.Int, error) {
	balance, err := conn.BalanceAt(context.Background(), common.HexToAddress(blockNumber), nil)
	if err != nil {
		log.Printf("balance error: %v\n", err)
		return nil, err
	}
	return balance, nil
}

// pass the transaction hash as Args
func GetTransactionReceipt(transactionHash string, conn *ethclient.Client) (*types.Receipt, error) {
	receipts, err := conn.TransactionReceipt(context.Background(), common.HexToHash(transactionHash))
	if err != nil {
		log.Printf("receipt error: %v\n", err)
		return nil, err
	}
	return receipts, nil
}

// pass in the transaction hash
// i abstracted the receipts func
func GetTransactions(ctx context.Context, transactionHash string, conn *ethclient.Client) {
	receipt, err := GetTransactionReceipt(transactionHash, conn)
	if err != nil {
		log.Printf("error fetching tx receipt: %v\n", err)
		return
	}
	for receipt != nil {
		// get block Number
		block, err := conn.BlockByNumber(ctx, receipt.BlockNumber)
		if err != nil {
			log.Printf("error fetching Block by number: %v\n", err)
			return
		}
		// iterate through the block data returned
		// and loop through the transactions
		// we can retrieve alot of data
		for _, transactions := range block.Transactions() {
			fmt.Printf("chain ID: %v", transactions.ChainId())
			fmt.Printf("receipient: %v\n", transactions.To())
			fmt.Printf("gasPrice: %v\n", transactions.GasPrice())
		}
	}
}
