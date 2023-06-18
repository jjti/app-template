package api

import (
	pb "github.com/jjtimmons/sblast/gen/server"
)

type service struct {
	pb.UnimplementedEchoServiceServer
}

var _ pb.EchoServiceServer = &service{}

// New returns a new gRPC servers' HTTP handler for use on the same address serving the UI.
func New() pb.EchoServiceServer {
	return new(service)
}
