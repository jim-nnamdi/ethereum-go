package custom

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SendEther(ctx context.Context, privkey string, debitAmount int64, recipient string, conn *ethclient.Client) error {
	// ethereum uses the Elliptic curve digital
	// signature algorithm
	privatekey, err := crypto.HexToECDSA(privkey)
	if err != nil {
		log.Println("privatekey error", err)
		return err
	}
	// get the publickey from private key
	pubkey := privatekey.Public()
	publickey, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		log.Print("error casting private key to public\n")
		return err
	}

	// get the address from the public key
	fromAddress := crypto.PubkeyToAddress(*publickey)
	nonce, err := conn.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Printf("error fetching nonce: %v\n", err)
		return err
	}

	// amount to send
	value := big.NewInt(debitAmount)
	gaslimit := uint64(21000)
	gasPrice, err := conn.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("no gas:%v\n", err)
		return err
	}

	chainID, _ := conn.ChainID(context.Background())
	gasTip, _ := conn.SuggestGasTipCap(context.Background())

	// address to send
	toAddress := common.HexToAddress(recipient)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTip,
		GasFeeCap: gasPrice,
		Gas:       gaslimit,
		To:        (*common.Address)(&toAddress),
		Value:     value,
		Data:      nil,
	})

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privatekey)
	if err != nil {
		log.Printf("cannot sign transaction: %v\n", err)
		return err
	}
	return conn.SendTransaction(context.Background(), signedTx)
}
