package consumer

import (
	"context"
	"fmt"
	"github.com/rafaeljesus/rabbus"
	"time"
)

type Request struct {
	Uri          string
	Exchange     string
	ExchangeType string
	Queue        string
	RoutingKey   string
	ConsumerTag  string
	Qos          int
}

func NewConsumer(request Request) (chan rabbus.ConsumerMessage, error) {
	r, err := rabbus.New(
		request.Uri,
		rabbus.Durable(true),
		rabbus.Attempts(5),
		rabbus.Sleep(time.Second*2),
		rabbus.Threshold(3),
		rabbus.PrefetchCount(request.Qos),
	)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go r.Run(ctx)

	return r.Listen(rabbus.ListenConfig{
		Exchange: request.Exchange,
		Kind:     request.ExchangeType,
		Key:      request.RoutingKey,
		Queue:    request.Queue,
	})
}
