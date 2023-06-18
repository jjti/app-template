package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// EchoStore is an interface for the Echo table.
type EchoStore interface {
	Create(ctx context.Context, echo Echo) error
}

type echoImpl struct {
	ddb   *dynamodb.Client
	table string
}

var _ EchoStore = echoImpl{}

// NewEchoStore returns a new EchoStore.
func NewEchoStore(ctx context.Context) (EchoStore, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &echoImpl{
		ddb:   dynamodb.NewFromConfig(sdkConfig),
		table: "echo-table",
	}, nil
}
