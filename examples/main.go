package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/emretiryaki/rabbitmq"
	"github.com/go-redis/redis"
	"gofx/consumer"
	"gofx/repository"
	"gofx/repository/clients"
	"time"
)

var (
	uri           = flag.Args()
	username      = flag.String("username", "guest", "username")
	password      = flag.String("password", "guest", "password")
	retryCount    = flag.Int("retry-count", 1, "retry count")
	prefetchCount = flag.Int("prefetch-count", 10, "prefetch count")
	exchange      = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType  = flag.Int("exchange-type", 3, "Exchange type -> 1-direct,2-fanout,3-topic,4-ConsistentHashing,5-XDelayedMessage")
	queue         = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
	bindingKey    = flag.String("key", "test-key", "AMQP binding key")
)

func init() {
	flag.Parse()
}

type (
	PersonV1 struct {
		Name    string
		Surname string
		City    City
		Count   int
	}

	City struct {
		Name string
	}
)

func main() {

	client := &repository.Client{
		Context: &clients.RedisClient{
			Conn: redis.NewClient(&redis.Options{}),
		},
	}

	onConsumed := func(message rabbitmq.Message) error {
		err := client.Add(message, 0)
		var consumeMessage PersonV1
		err = json.Unmarshal(message.Payload, &consumeMessage)
		client.Add(consumeMessage, time.Duration(0))
		if err != nil {
			return err
		}
		fmt.Println(time.Now().Format("Mon, 02 Jan 2006 15:04:05 "), " Message:", consumeMessage)

		return nil
	}

	r, c := consumer.AddConsumer(consumer.Request{
		Uri:           []string{"127.0.0.1:5672", "127.0.0.1:5673", "127.0.0.1:5672"},
		UserName:      *username,
		Password:      *password,
		RetryCount:    *retryCount,
		PrefetchCount: *prefetchCount,
		Exchange:      *exchange,
		ExchangeType:  *exchangeType,
		Queue:         *queue,
		RoutingKey:    *bindingKey,
	})

	c.HandleConsumer(onConsumed)
	r.RunConsumers()
}
