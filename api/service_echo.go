package api

import (
	"context"
	"fmt"

	pb "github.com/jjtimmons/sblast/gen/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/go-hclog"
)

func (s *service) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log := hclog.FromContext(ctx)

	// validate request
	if err := req.ValidateAll(); err != nil {
		log.Error("received invalid request", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info("received echo request", "input", req.GetInput())
	return &pb.EchoResponse{
		Output: fmt.Sprintf("echo: %v", req.GetInput()),
	}, nil
}
