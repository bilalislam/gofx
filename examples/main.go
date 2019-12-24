package main

import (
	"flag"
	"fmt"
	"gofx/consumer"
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
	qos          = flag.Int("qos", 10, "qos")
)

func init() {
	flag.Parse()
}

func main() {
	messages, err := consumer.NewConsumer(consumer.Request{
		Uri:          *uri,
		Exchange:     *exchange,
		ExchangeType: *exchangeType,
		Queue:        *queue,
		RoutingKey:   *bindingKey,
		ConsumerTag:  *consumerTag,
		Qos:          *qos,
	})

	if err != nil {
		fmt.Errorf("%s", err)
	}

	if err != nil {
		log.Fatalf("Failed to create listener %s", err)
	}

	defer close(messages)

	for {
		log.Println("Listening for messages...")

		m, ok := <-messages
		if !ok {
			log.Println("Stop listening messages!")
		}

		m.Ack(false)
		log.Println("Message was consumed")
	}
}
