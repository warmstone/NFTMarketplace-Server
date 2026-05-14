package client

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient(ctx context.Context, rpcURL string) (*ethclient.Client, error) {
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}
