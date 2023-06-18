package api

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/jjtimmons/sblast/pb/server"
	"github.com/jjtimmons/sblast/pkg/db"
)

func (s *service) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log := hclog.FromContext(ctx)

	// validate request
	if err := req.ValidateAll(); err != nil {
		log.Error("received invalid request", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// write the echo to the db
	id := uuid.NewString()
	if err := s.store.Create(ctx, db.Echo{
		ID:    id,
		Value: req.GetInput(),
	}); err != nil {
		log.Error("failed to insert echo to the database", "error", err)
		return nil, status.Error(codes.Internal, "failed to add echo to store")
	}

	log.Info("received echo request", "input", req.GetInput(), "id", id)
	return &pb.EchoResponse{
		Output: fmt.Sprintf("echo: %v", req.GetInput()),
	}, nil
}
