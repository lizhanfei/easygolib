package grpc

import (
	"context"
	"testing"
)

func TestGrpc(t *testing.T) {
	address := ":19090"
	ctxBase := context.Background()
	grpcServer := NewServer(ctxBase)

	RegisterLibTestServiceServer(grpcServer.Server, &LibTestService{})
	lis, err := NewListenerTcp(address)
	if err != nil {
		panic(err)
	}
	grpcServer.SetListener(lis)
	//grpcServer.Run()
}

type LibTestService struct {
}

func (this *LibTestService) LibAuth(ctx context.Context, req *LibAuthRequest) (*LibAuthResponse, error) {
	return &LibAuthResponse{
		Id:     0,
		Name:   "authRes_" + req.Token,
		Status: 1,
	}, nil
}
