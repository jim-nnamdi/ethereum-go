package custom

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ChainConnect() (*ethclient.Client, error) {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/5c0531336045410e9bb1c1e0d3fec3eb")
	if err != nil {
		log.Printf("err connecting to ethnode:%v\n", err)
		return nil, err
	}
	return conn, nil
}
