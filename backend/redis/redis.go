package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"tcms/m/dry"
)

type Client interface {
	Publish(ctx context.Context, channel string, payload interface{}) (*redis.IntCmd, error)
	Subscribe(ctx context.Context, channel string) *redis.PubSub
}

type client struct {
	client *redis.Client
}

func (c client) Publish(ctx context.Context, channel string, payload interface{}) (*redis.IntCmd, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.client.Publish(ctx, channel, data), nil
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
