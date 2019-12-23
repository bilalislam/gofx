
How can I write a new consumer?
---------------------------

* Import `"github.com/bilalislam/torc/consumer"` package.
* Create a new instance with `consumer.NewConsumer()`.
* Start consume all messages
* Call `c.Consume()`


For print help for example
---------------------------

```sh
$ go build consumer -o consumer
$ ./consumer --help
```

Simple Usage 
---------------------------

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
Docker file
---------------------------

```dockerfile
# Dockerfile Example
# https://medium.com/@petomalina/using-go-mod-download-to-speed-up-golang-docker-builds-707591336888
# Based on this image: https:/hub.docker.com/_/golang/
FROM golang:latest as builder


RUN mkdir -p /go/src/github.com/bilalislam/torc
WORKDIR /go/src/github.com/bilalislam/torc

# Force the go compiler to use modules
ENV GO111MODULE on
# <- COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .


# Compile application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o simle-consumer

RUN chmod +x /go/src/github.com/bilalislam/torc

#Image Diff
#(Not Scratch) 1.23GB
#(Scratch    ) 34.3MB
# <- Second step to build minimal image
FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/bilalislam/torc .
ENV APP_ENV qa
# Execite application when container is started
CMD ["./simple-consumer"]


EXPOSE 8080

```