package handler

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	e.GET("/health-check", healthCheck)
	e.GET("/users/:id/basket", handler.getBasketById)
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
