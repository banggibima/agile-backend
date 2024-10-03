package redis

import (
	"context"

	"github.com/banggibima/agile-backend/config"
	"github.com/redis/go-redis/v9"
)

func Client(config *config.Config) (*redis.Client, error) {
	url := config.Redis.Url

	options, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)

	if err := Connect(client); err != nil {
		return nil, err
	}

	return client, nil
}

func Connect(client *redis.Client) error {
	ctx := context.Background()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}
