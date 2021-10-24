package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"tcms/m/dry"
)

type Client interface {
	Publish(ctx context.Context, channel string, payload interface{})
	Subscribe(ctx context.Context, channel string) *redis.PubSub
}

type client struct {
	client *redis.Client
}

func (c client) Publish(ctx context.Context, channel string, payload interface{}) {
	c.client.Publish(ctx, channel, payload)
}

func (c client) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return c.client.Subscribe(ctx, channel)
}

func GetClient() Client {

	host, err := getRedisHost()
	dry.HandleErrorPanic(err)

	password := getRedisPassword()

	db, err := getRedisDatabase()
	dry.HandleError(err)

	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	return &client{client: rdb}
}
