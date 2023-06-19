package redis_event

import (
	"fmt"
	"time"
)

func (event *Event) GetAll() (map[string]string, error) {
	ctx, cancel := event.ctx()
	defer cancel()

	allKeys, err := event.connection.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	ctx, cancel = event.ctx(time.Second * 5)
	defer cancel()

	results := make(map[string]string)
	for _, key := range allKeys {
		valueTypeObject := event.connection.Get(ctx, key)
		if valueTypeObject.Err() != nil {
			return nil, err
		}

		value, err := valueTypeObject.Result()
		if err != nil {
			fmt.Println("get result:", err)
			return nil, err
		}

		fmt.Println("value:", value)
		results[key] = value
	}

	return results, nil
}

func (event *Event) Push(queue string, values ...any) error {
	ctx, cancel := event.ctx()
	defer cancel()

	return event.connection.RPush(ctx, queue, values...).Err()
}

func (event *Event) Consume(queue string) chan any {
	go event.consume(queue)
	return nil
}

func (event *Event) GetByKey(key string) (string, error) {
	ctx, cancel := event.ctx()
	defer cancel()

	result, err := event.connection.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (event *Event) Set(key string, value any, expiration ...time.Duration) error {
	ctx, cancel := event.ctx()
	defer cancel()

	var expirationTime time.Duration
	if len(expiration) > 0 {
		expirationTime = expiration[0]
	}

	return event.connection.Set(ctx, key, value, expirationTime).Err()
}

func (event *Event) Increment(key string) error {
	ctx, cancel := event.ctx()
	defer cancel()
	return event.connection.Incr(ctx, key).Err()
}

func (event *Event) Decrement(key string) error {
	ctx, cancel := event.ctx()
	defer cancel()
	return event.connection.Decr(ctx, key).Err()
}
