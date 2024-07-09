package redis

import (
	"context"
	"fmt"
)

func InitClient() *Client {
	network := "tcp"
	address := "127.0.0.1:6379"
	password := ""

	client := NewRedis(network, address, password)
	conn, err := client.pool.GetContext(context.Background())
	if err != nil {
		panic(err)
	}
	_, err1 := conn.Do("PING")
	if err1 != nil {
		panic(fmt.Errorf("redis fail to connect: %v", err1))
	}
	return client

}
