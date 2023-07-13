package handlers

import (
	"context"

	pb "github.com/julkhong/blockchain-server"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.BlockchainServer {
	return blockchainService{}
}

type blockchainService struct{}

func (s blockchainService) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	return &resp, nil
}
