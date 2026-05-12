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

func NewWsClient(ctx context.Context, wsURL string) (*ethclient.Client, error) {
	wsClient, err := ethclient.DialContext(ctx, wsURL)
	if err != nil {
		return nil, err
	}

	return wsClient, nil
}
