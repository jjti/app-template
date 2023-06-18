package api

import (
	pb "github.com/jjtimmons/sblast/gen/sblast"
)

type service struct {
	pb.UnimplementedSemanticBlastServiceServer
}

var _ pb.SemanticBlastServiceServer = &service{}

// New returns a new gRPC servers' HTTP handler for use on the same address serving the UI.
func New() pb.SemanticBlastServiceServer {
	return new(service)
}
