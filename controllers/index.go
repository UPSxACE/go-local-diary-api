package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetIndexRoutes(e *echo.Echo) {
	e.GET("/", GetIndexController)
	e.POST("/htmx/change-message", HtmxChangeMessage)
}

func GetIndexController(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"Message": "Hello World!",
	})
}

type ChangeMessageRequest struct {
	Name string `json:"name" form:"name" query:"name"`
}

func HtmxChangeMessage(c echo.Context) error {
	request := new(ChangeMessageRequest)
	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	time.Sleep(1 * time.Second)
	return c.Render(http.StatusOK, "updated-message", map[string]interface{}{
		"Message": fmt.Sprintf("Hello %v!", request.Name),
	})
}
