// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: a2b01cac16
// Version Date: 2022-10-20T18:44:52Z

// Package grpc provides a gRPC client for the Blockchain service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/julkhong/blockchain-server"
	"github.com/julkhong/blockchain-server/blockchain-service/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.BlockchainServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var echoEndpoint endpoint.Endpoint
	{
		echoEndpoint = grpctransport.NewClient(
			conn,
			"blockchain.Blockchain",
			"Echo",
			EncodeGRPCEchoRequest,
			DecodeGRPCEchoResponse,
			pb.EchoResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		EchoEndpoint: echoEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCEchoResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC echo reply to a user-domain echo response. Primarily useful in a client.
func DecodeGRPCEchoResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.EchoResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCEchoRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain echo request to a gRPC echo request. Primarily useful in a client.
func EncodeGRPCEchoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.EchoRequest)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
