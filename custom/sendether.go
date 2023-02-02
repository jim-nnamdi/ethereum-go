package custom

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendEther(ctx context.Context, privkey string, debitAmount int64, recipient string, conn *ethclient.Client) {
	// ethereum uses the Elliptic curve digital
	// signature algorithm
	privatekey, err := crypto.HexToECDSA(privkey)
	if err != nil {
		log.Println("privatekey error", err)
		return
	}
	// get the publickey from private key
	pubkey := privatekey.Public()
	publickey, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		log.Printf("error casting private key to public: %v\n")
		return
	}

	// get the address from the public key
	fromAddress := crypto.PubkeyToAddress(*publickey)
	nonce, err := conn.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Printf("error fetching nonce: %v\n", err)
		return
	}

	// amount to send
	value := big.NewInt(debitAmount)
	gaslimit := uint64(21000)
	gasPrice, err := conn.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("no gas:%v\n", err)
		return
	}

	// address to send
	toAddress := common.HexToAddress(recipient)

	// WIP
}
