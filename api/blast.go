package api

import (
	"context"

	"github.com/biogo/ncbi/blast"
	pb "github.com/jjtimmons/sblast/gen/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/go-hclog"
)

func (s *service) Blast(ctx context.Context, req *pb.BlastRequest) (*pb.BlastResponse, error) {
	log := hclog.Default()

	// validate request
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// issue a request to the BLAST server
	rid, err := blast.Put(req.Seq, &blast.PutParameters{
		NuclPenalty: -3,
		NuclReward:  2,
		GapCosts:    [2]int{5, 2},
		Database:    "nt",
		Program:     "blastn",
		WordSize:    11,
	}, "sblast", "joshuatimmons1@gmail.com")
	if err != nil {
		log.Error("failed to start a blast request", "error", err)
		return nil, status.Error(codes.Internal, "failed to start the blast request")
	}
	log.Info("successfully started blast request", "rid", rid.String())

	return &pb.BlastResponse{
		Rid: "abc",
	}, nil
}
