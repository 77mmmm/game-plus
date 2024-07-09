package redis

import (
	"context"
	"errors"
	"github.com/gomodule/redigo/redis"
	"strings"
	"time"
)

type Client struct {
	pool *redis.Pool
	ClientOptions
}

type BoardClient interface {
	ZAdd(ctx context.Context, key string, username int64, userscore int64) (int64, error)
	ZRank(ctx context.Context, key string, username int64) (int64, error)
	ZScore(ctx context.Context, key string, username int64) (int64, error)
	ZCount(ctx context.Context, key string) (int64, error)
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)
}

func NewRedis(network, address, password string, opts ...ClientOption) *Client {
	c := &Client{
		ClientOptions: ClientOptions{
			address:  address,
			password: password,
			network:  network,
		},
	}
	for _, opt := range opts {
		opt(&c.ClientOptions)
	}
	repairClient(&c.ClientOptions)
	pool := c.getPool()
	return &Client{
		pool: pool,
	}

}

func (c *Client) getPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     c.maxIdle,
		IdleTimeout: time.Duration(c.idleTimeoutSeconds) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := c.getConn()
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxActive: c.maxActive,
		Wait:      c.wait,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (c *Client) getConn() (redis.Conn, error) {
	if c.address == "" {
		panic("cannot get redis address from config")
	}
	var diaOpts []redis.DialOption
	if len(c.password) > 0 {
		diaOpts = append(diaOpts, redis.DialPassword(c.password))
	}
	conn, err := redis.Dial(c.network, c.address, diaOpts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) ZAdd(ctx context.Context, key string, username int64, userscore int64) (int64, error) {
	if key == "" {
		return -1, errors.New("redis ZAdd key can't be empty")
	}
	if userscore <= 0 {
		return -1, errors.New("userscore cannot be less than 0")
	}
	conn, err := c.pool.GetContext(ctx)
	defer conn.Close()
	if err != nil {
		return -1, err
	}
	reply, err1 := conn.Do("ZADD", key, username, userscore)
	if err1 != nil {
		return -1, err1
	}
	if resp, ok := reply.(string); ok && strings.ToLower(resp) == "ok" {
		return 1, nil
	}
	return reply.(int64), err1

}

func (c *Client) ZRank(ctx context.Context, key string, username int64) (int64, error) {
	if key == "" {
		return -1, errors.New("redis ZAdd key can't be empty")
	}
	conn, err := c.pool.GetContext(ctx)
	defer conn.Close()
	if err != nil {
		return -1, err
	}
	reply, err1 := conn.Do("ZRANK", key, username)
	if err1 != nil || reply == -1 {
		return -1, err1
	}
	return reply.(int64), nil
}

func (c *Client) ZScore(ctx context.Context, key string, username int64) (int64, error) {
	conn, err := c.pool.GetContext(ctx)
	defer conn.Close()
	if err != nil {
		return -1, err
	}
	reply, err1 := conn.Do("ZSCORE", key, username)
	if err1 != nil {
		return -1, err1
	}

	return reply.(int64), nil

}

func (c *Client) ZCount(ctx context.Context, key string) (int64, error) {
	if key == "" {
		return -1, errors.New("redis ZCount key can't be empty")
	}
	conn, err := c.pool.GetContext(ctx)
	defer conn.Close()
	if err != nil {
		return -1, err
	}
	reply, err1 := conn.Do("ZCARD", key)
	if err1 != nil {
		return -1, err1
	}
	return reply.(int64), nil
}

func (c *Client) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	if key == "" {
		return nil, errors.New("redis Zrange key can't be empty")
	}
	conn, err := c.pool.GetContext(ctx)
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	withscores := "WITHSCORES"
	reply, err1 := conn.Do("RANGE", key, start, stop, withscores)
	if err1 != nil {
		return nil, err1
	}
	return reply.([]string), nil
}
