package main

import (
	"flag"
	"github.com/go-redis/redis"
	"gofx/consumer"
	"gofx/repository"
	"gofx/repository/clients"
	"log"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchange     = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	queue        = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
	bindingKey   = flag.String("key", "test-key", "AMQP binding key")
	consumerTag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
	environment  = flag.String("environment", "dev", "os environment")
)

func init() {
	flag.Parse()
}

func main() {

	c, err := consumer.NewConsumer(consumer.Request{
		Uri:          *uri,
		Exchange:     *exchange,
		ExchangeType: *exchangeType,
		Queue:        *queue,
		RoutingKey:   *bindingKey,
		ConsumerTag:  *consumerTag,
	})

	if err != nil {
		log.Fatalf("%s", err)
	}

	messages, err := c.Consume(*queue)
	for d := range messages {
		log.Printf(
			"got %dB delivery: [%v] %s",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)

		client := &repository.Client{
			Context: &clients.RedisClient{
				Conn: redis.NewClient(&redis.Options{}),
			},
		}

		err := client.Add(d.Body, 0)
		if err == nil {
			_ = d.Ack(true)
		} else {
			panic(err)
		}
	}
}
