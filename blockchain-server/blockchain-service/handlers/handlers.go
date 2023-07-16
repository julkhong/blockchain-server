package handlers

import (
	"context"
	"fmt"
	pb "github.com/julkhong/blockchain-server"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.BlockchainServer {
	//// Initialize the Fabric SDK
	//sdk := initializeSDK()
	//defer sdk.Close()
	//
	//// Create a client and context to interact with the smart contract
	////TODO move this out to a service config file
	//channelID := "mychannel"
	//chaincodeID := "mychaincode"
	//username := "user1"
	//client, ctx := createClient(sdk, channelID, chaincodeID, username)
	return blockchainService{
		//client,
		//ctx,
		//channelID,
	}
}

type blockchainService struct {
	//ChannelClient *channel.Client
	//ClientContext contextAPI.ClientProvider
	//ChannelID     string
}

func (s blockchainService) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	var resp pb.EchoResponse
	resp.Out = in.In
	return &resp, nil
}

func (s blockchainService) Send(ctx context.Context, in *pb.ChainCodeRequest) (*pb.ChainCodeResponse, error) {
	var resp pb.ChainCodeResponse
	req := in
	if req.Params.ReceiverID == "" {
		log.WithFields(log.Fields{"request": req}).Error("Empty receiver id")
		return nil, errors.New("Bad request from client")
	}
	if req.Params.SenderID == "" {
		log.WithFields(log.Fields{"request": req}).Error("Empty sender id")
		return nil, errors.New("Bad request from client")
	}
	if !validateChainCodeRequest(req) {
		log.WithFields(log.Fields{"request": req}).Error("Request is not valid")
		return nil, errors.New("Bad request from client")
	}
	amount := uint64(math.Round(float64(req.Params.Amount)))
	fmt.Println(amount)
	//logic.Send(s.ChannelClient, s.ClientContext, req.Params.SenderID, req.Params.ReceiverID, amount, s.ChannelID)
	return &resp, nil
}

func (s blockchainService) Balance(ctx context.Context, in *pb.ChainCodeRequest) (*pb.ChainCodeResponse, error) {
	var resp pb.ChainCodeResponse
	req := in
	if req.Params.AccountID == "" {
		log.WithFields(log.Fields{"request": req}).Error("Empty account id")
		return nil, errors.New("Bad request from client")
	}
	if !validateChainCodeRequest(req) {
		log.WithFields(log.Fields{"request": req}).Error("Request is not valid")
		return nil, errors.New("Bad request from client")
	}
	return &resp, nil
}
