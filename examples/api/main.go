package main

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gofx/examples/api/handler"
	"gofx/repository"
	"gofx/repository/clients"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
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
