package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	/*v1 := e.Group("/api/v1")
	{
		baskets := v1.Group("/baskets")
		{
			baskets.GET(":id", handler.getBasketById)
		}
	}*/

	e.GET("/health-check", healthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/baskets/:id", handler.getBasketById)

}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healhly !")
}

// ShowAccount godoc
// @Summary Show a basket
// @Description get string by ID
// @Tags baskets
// @Accept  json
// @Produce  json
// @Param id path int true "Customer ID"
// @Router /baskets/{id} [get]
// @Success 200 {object} HTTPOk
// @Failure 400 {object} HTTPError
func (h *Handler) getBasketById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return NewError(c, http.StatusBadRequest, errors.New("id not found"))
	}
	_ = h.repository.Get(id)

	return NewSucces(c, http.StatusOK, id)
}

// NewError example
func NewError(ctx echo.Context, status int, err error) error {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)

	return err
}

// NewSucces example
func NewSucces(ctx echo.Context, status int, result interface{}) error {
	return ctx.JSON(status, result)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPError example
type HTTPOk struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"status ok request"`
}
