package custom

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestSendEther(t *testing.T) {
	type Args struct {
		ctx       context.Context
		HexKey    string
		Recipient string
	}

	tests := []struct {
		name             string
		args             Args
		expectedResponse error
	}{
		{
			name: "sending ether successfully",
			args: Args{
				ctx:       context.Background(),
				HexKey:    "fedbec12b2bc4ecf0c965c64003a9ee384ad59a8831a7f2e783d7d4680aca18a",
				Recipient: "0x37d32eDC4F12c058d3Ede9c3a5D253F08eBB6ce9",
			},
			expectedResponse: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			conn, err := ethclient.Dial("/Users/mac/Library/Ethereum/geth.ipc")
			if err != nil {
				log.Printf("error! failed to connect: %v\n", err)
				x := fmt.Errorf("%v\n", err)
				fmt.Print(x)
			}
			got := ExperimentalSendEther(tc.args.ctx, tc.args.HexKey, tc.args.Recipient, conn)
			assert.Equal(t, tc.expectedResponse, got)
		})
	}
}
