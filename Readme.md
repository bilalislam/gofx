
How to package
---------------
```sh
$ go mod init gofx
$ go mod tidy
$ go build
```

You can use with below;
```go
import "gofx/consumer"
```

Consumer
---------------
```go
func main() {
	c, err := consumer.NewConsumer(*uri, *exchange, *exchangeType, *queue, *bindingKey, *consumerTag)
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
	}
}
```


Bridge Pattern for storege
------------------------------

```go
client := &repository.Client{
			Context: &clients.RedisClient{
				Conn: redis.NewClient(&redisOptions),
			},
		}

		err := client.Add(d.Body, 0)
```

#####references
* https://github.com/tmrts/go-patterns
* http://blog.ralch.com/tutorial/design-patterns/golang-bridge/