// Copyright 2020 Jim Nnamdi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

	// Firstly we need a private key of the sender
	// Which would help us generate a public key
	// and corresponding public address.
	privatekey, err := crypto.HexToECDSA(privkey)
	if err != nil {
		log.Println("privatekey error", err)
		return err
	}

	// Generate the public key from the privatekey
	// The private and public key work together
	// Hash a transaction using the public key
	// and read the information using the private.
	pubkey := privatekey.Public()
	publickey, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		log.Print("error casting private key to public\n")
		return err
	}

	// Get the address of the sender from the publickey
	// This would be the address to which the Ether would
	// be deducted from to be sent to the recipient
	fromAddress := crypto.PubkeyToAddress(*publickey)
	nonce, err := conn.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Printf("error fetching nonce: %v\n", err)
		return err
	}

	// Set other generic data values such as the amount
	// and the gaslimit and the gasprice of the transaction
	value := big.NewInt(debitAmount)
	gaslimit := uint64(21000)
	gasPrice, err := conn.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("no gas:%v\n", err)
		return err
	}

	chainID, _ := conn.ChainID(context.Background())
	gasTip, _ := conn.SuggestGasTipCap(context.Background())
	toAddress := common.HexToAddress(recipient)

	// initiate the transaction to be done using the
	// NewTx interface that can be satisfied by the
	// DynamicFeeTx and the AccessListTx
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

	// Sign the transaction finally using the NewLondonSigner
	// to verify the transaction carried out on the initiated
	// transactions.
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privatekey)
	if err != nil {
		log.Printf("cannot sign transaction: %v\n", err)
		return err
	}
	return conn.SendTransaction(context.Background(), signedTx)
}

func experimentalSendEther(ctx context.Context, expHexKey string, recipient string, conn *ethclient.Client) error {
	experimentalPrivateKey, err := crypto.HexToECDSA(expHexKey)
	if err != nil {
		log.Printf("bad private key: %v\n", err)
		return err
	}
	experimentalPublicKey := experimentalPrivateKey.Public()
	realExperimentalPubKey := experimentalPublicKey.(*ecdsa.PublicKey)
	realExperimentalAddress := crypto.PubkeyToAddress(*realExperimentalPubKey)

	experimentalNonce, err := conn.PendingNonceAt(ctx, realExperimentalAddress)
	if err != nil {
		log.Printf("bad nonce val: %v\n", err)
		return err
	}
	experimentalRecipient := common.HexToAddress(recipient)
	experimentalChainId, _ := conn.ChainID(ctx)
	experimentalGasPrice := big.NewInt(5000000000000000000)
	experimentalValue, _ := conn.SuggestGasPrice(ctx)
	experimentalGasLimit := uint64(21000)

	experimentalTx := types.NewTx(&types.AccessListTx{
		ChainID:  experimentalChainId,
		Nonce:    experimentalNonce,
		GasPrice: experimentalGasPrice,
		Gas:      experimentalGasLimit,
		To:       (*common.Address)(&experimentalRecipient),
		Value:    experimentalValue,
		Data:     nil,
	})

	experimentSignedTx, err := types.SignTx(experimentalTx, types.NewLondonSigner(experimentalChainId), experimentalPrivateKey)
	if err != nil {
		log.Printf("error signing transaction: %v\n", err)
		return err
	}
	return conn.SendTransaction(ctx, experimentSignedTx)
}
