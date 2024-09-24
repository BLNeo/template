package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
)

type Config struct {
	DBNumber    int      `yaml:"dbNumber"`
	Addresses   []string `yaml:"addresses"`
	Password    string   `yaml:"password"`
	MaxIdle     int      `yaml:"maxIdle"`
	MaxActive   int      `yaml:"maxActive"`
	IdleTimeout int      `yaml:"idleTimeout"`
}

var (
	once sync.Once
	rdb  *redis.Client
)

func GetRdb() *redis.Client {
	return rdb
}

// InitClient 初始化一个ENGIN
func InitClient(in *Config) error {
	if len(in.Addresses) == 0 {
		return errors.New("addresses is empty")
	}

	client := redis.NewClient(&redis.Options{
		Addr:           in.Addresses[0],
		Password:       in.Password,
		DB:             in.DBNumber,
		MaxActiveConns: in.MaxActive,
		MaxIdleConns:   in.MaxIdle,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	once.Do(func() {
		rdb = client
	})
	fmt.Println("使用redis单节点方式")
	return nil
}
