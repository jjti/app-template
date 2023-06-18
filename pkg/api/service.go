package api

import (
	"context"
	"fmt"

	pb "github.com/jjtimmons/sblast/pb/server"
	"github.com/jjtimmons/sblast/pkg/db"
)

type service struct {
	pb.UnimplementedEchoServiceServer

	store db.EchoStore
}

var _ pb.EchoServiceServer = &service{}

// New returns a new gRPC servers' HTTP handler for use on the same address serving the UI.
func New(ctx context.Context) (pb.EchoServiceServer, error) {
	store, err := db.NewEchoStore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create store: %w", err)
	}

	return &service{
		store: store,
	}, nil
}
