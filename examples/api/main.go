package main

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gofx/examples/api/handler"
	"gofx/repository"
	"gofx/repository/clients"
)

func main() {
	e := echo.New()

	client := &repository.Client{
		Context: &clients.RedisClient{
			Conn: redis.NewClient(&redis.Options{}),
		},
	}

	handler.NewHandler(e, *client)
	e.Logger.Fatal(e.Start(":1323"))
}
