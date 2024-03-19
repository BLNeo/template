package service

import (
	"context"
	signPb "github.com/BLNeo/protobuf-grpc-file/sign"
	"sync"
	"time"
)

var (
	signRpcClient signPb.SignClient
	once          sync.Once
)

func InitSignRpcClient(client signPb.SignClient) {
	once.Do(func() {
		signRpcClient = client
	})
}

func NewSignRpcService() ISignRpc {
	return &SignRpcService{
		signRpcClient: signRpcClient,
	}
}

type ISignRpc interface {
	VerifyToken(token string) error
}

type SignRpcService struct {
	signRpcClient signPb.SignClient
}

func (s *SignRpcService) VerifyToken(token string) error {
	ctx, f := context.WithTimeout(context.Background(), time.Duration(3*int(time.Second)))
	defer f()
	in := &signPb.VerifyTokenRequest{
		Token: token,
	}
	_, err := s.signRpcClient.VerifyToken(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
