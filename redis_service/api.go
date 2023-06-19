package redis_service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewConnection(address, password string, database int) (*redis.Client, error) {
	connection := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       database,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := connection.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return connection, nil
}
