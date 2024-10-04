package redisclient

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	pingTimeout = 5 * time.Second
)

type Config struct {
	Address string
}

type Client struct {
	Conn *redis.Client
}

func New(cfg *Config) *Client {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr: cfg.Address,
	})

	status := client.Ping(ctx)
	err := status.Err()

	if err != nil {
		panic(err)
	}

	return &Client{
		Conn: client,
	}
}

func (c *Client) Close() error {
	return c.Conn.Close()
}
