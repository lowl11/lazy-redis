package redis_event

import (
	"github.com/lowl11/lazy-redis/redis_service"
	"github.com/redis/go-redis/v9"
)

type Event struct {
	connection     *redis.Client
	consumeChannel chan string
}

func New(address, password string, database int) (*Event, error) {
	connection, err := redis_service.NewConnection(address, password, database)
	if err != nil {
		return nil, err
	}

	return &Event{
		connection:     connection,
		consumeChannel: make(chan string),
	}, nil
}
