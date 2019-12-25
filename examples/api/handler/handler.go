package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	_ "gofx/examples/api/docs"
	"gofx/repository"
	"net/http"
)

type Handler struct {
	repository repository.Client
}

func NewHandler(e *echo.Echo, client repository.Client) {

	handler := &Handler{
		repository: client,
	}

	v1 := e.Group("/api/v1")
	{
		baskets := v1.Group("/basket")
		{
			baskets.GET(":id", handler.getBasketById)
		}
	}

	e.GET("/health-check", healthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healhly !")
}

func (h *Handler) getBasketById(c echo.Context) error {
	id := c.Param("id")
	_ = h.repository.Get(id)

	return c.JSON(http.StatusOK, id)
}
