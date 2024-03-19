package service

import (
	signPb "github.com/BLNeo/protobuf-grpc-file/sign"
	"sync"
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

func GetSignRpcClient() signPb.SignClient {
	return signRpcClient
}
