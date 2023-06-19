package redis_event

import (
	"context"
	"log"
	"time"
)

func (event *Event) ctx(customTimeout ...time.Duration) (context.Context, func()) {
	deadlineTimeout := time.Second * 3
	if len(customTimeout) > 0 {
		deadlineTimeout = customTimeout[0]
	}
	return context.WithTimeout(context.Background(), deadlineTimeout)
}

func (event *Event) consume(key string) {
	for {
		if err := event.receiveValue(key); err != nil {
			log.Println(err)
		}
	}
}

func (event *Event) receiveValue(key string) error {
	ctx, cancel := event.ctx()
	defer cancel()

	item, err := event.connection.LPop(ctx, key).Result()
	if err != nil {
		return err
	}

	event.consumeChannel <- item
	return nil
}
