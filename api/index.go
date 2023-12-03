package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"Message": "Hello World!",
	})
}

type ChangeMessageRequest struct {
	Name string `json:"name" form:"name" query:"name"`
}

func ChangeMessage(c echo.Context) error {
	request := new(ChangeMessageRequest)
	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	time.Sleep(1 * time.Second)
	return c.Render(http.StatusOK, "updated-message", map[string]interface{}{
		"Message": fmt.Sprintf("Hello %v!", request.Name),
	})
}
