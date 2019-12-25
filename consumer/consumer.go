package consumer

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Request struct {
	Uri          string
	Exchange     string
	ExchangeType string
	Queue        string
	RoutingKey   string
	ConsumerTag  string
}

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan string
}

func NewConsumer(request Request) (*Consumer, error) {
	c := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     request.ConsumerTag,
		done:    make(chan string),
	}

	var err error

	log.Printf("dialing %s", request.Uri)
	c.conn, err = amqp.Dial(request.Uri)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	log.Printf("got Channel, declaring Exchange (%s)", request.Exchange)
	if err = c.channel.ExchangeDeclare(
		request.Exchange,     // name of the exchange
		request.ExchangeType, // type
		false,                // durable
		false,                // delete when complete
		false,                // internal
		false,                // noWait
		nil,                  // arguments
	); err != nil {
		return nil, fmt.Errorf("Exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, declaring Queue (%s)", request.Queue)
	state, err := c.channel.QueueDeclare(
		request.Queue, // name of the queue
		true,          // durable
		false,         // delete when usused
		false,         // exclusive
		false,         // noWait
		nil,           // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	log.Printf("declared Queue (%d messages, %d consumers), binding to Exchange (key '%s')",
		state.Messages, state.Consumers, request.RoutingKey)

	if err = c.channel.QueueBind(
		request.Queue,      // name of the queue
		request.RoutingKey, // bindingKey
		request.Exchange,   // sourceExchange
		false,              // noWait
		nil,                // arguments
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}

	c.Qos(10,0,false)

	return c, nil
}

func (c *Consumer) Consume(queue string) (<-chan amqp.Delivery, error) {
	return c.channel.Consume(
		queue, // name
		c.tag, // consumerTag,
		false, // noAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // arguments
	)
}

func (c *Consumer) Qos(prefetchCount, prefetchSize int, global bool) {
	c.channel.Qos(prefetchCount, prefetchSize, global)
}
