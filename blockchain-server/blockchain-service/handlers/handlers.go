package handlers

import (
	"context"
	pb "github.com/julkhong/blockchain/blockchain-server"
	"math"
	"strconv"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.BlockchainServer {
	return blockchainService{}
}

type blockchainService struct {
}

func (s blockchainService) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	resp.Out = in.In
	return &resp, nil
}

func (s blockchainService) SendBalance(ctx context.Context, in *pb.ChainCodeRequest) (*pb.ChainCodeResponse, error) {
	ContractSetup()
	InitAccounts(&InvokedContract)
	var resp pb.ChainCodeResponse
	req := in
	//if req.Params.ReceiverID == "" {
	//	log.WithFields(log.Fields{"request": req}).Error("Empty receiver id")
	//	return nil, errors.New("Bad request from client")
	//}
	//if req.Params.SenderID == "" {
	//	log.WithFields(log.Fields{"request": req}).Error("Empty sender id")
	//	return nil, errors.New("Bad request from client")
	//}
	//if !validateChainCodeRequest(req) {
	//	log.WithFields(log.Fields{"request": req}).Error("Request is not valid")
	//	return nil, errors.New("Bad request from client")
	//}
	amount := uint64(math.Round(float64(req.Params.Amount)))
	Send(&InvokedContract, req.Params.SenderID, req.Params.ReceiverID, strconv.FormatUint(amount, 10))
	resp.Id = req.Id
	resp.Code = "ok"
	return &resp, nil
}

func (s blockchainService) GetBalance(ctx context.Context, in *pb.ChainCodeRequest) (*pb.ChainCodeResponse, error) {
	ContractSetup()
	InitAccounts(&InvokedContract)
	var resp pb.ChainCodeResponse
	req := in
	//if req.Params.AccountID == "" {
	//	log.WithFields(log.Fields{"request": req}).Error("Empty account id")
	//	return nil, errors.New("Bad request from client")
	//}
	//if !validateChainCodeRequest(req) {
	//	log.WithFields(log.Fields{"request": req}).Error("Request is not valid")
	//	return nil, errors.New("Bad request from client")
	//}
	amount := GetBalance(&InvokedContract, req.Params.AccountID)
	resp.Id = req.Id
	resp.Code = "ok"
	f, _ := strconv.ParseFloat(amount, 32)
	resp.Amount = float32(f)

	return &resp, nil
}
