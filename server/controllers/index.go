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

// GetIndex godoc
// @Summary Home page
// @Description Renders home page
// @Tags / route
// @Produce html
// @Router / [get]
// @Success 200 "Rendered page successefully"
func GetIndexController(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"Message": "Hello World!",
	})
}

type ChangeMessageRequest struct {
	Name string `json:"name" form:"name" query:"name"`
}

// HtmxChangeMessage godoc 
// @Summary Change message 
// @Description Updates message in home page 
// @Tags / route 
// @Produce html 
// @Router /htmx/change-message [post] 
// @Param name formData string true "Name that will be in the new message" 
// @Success 200 "Updated message successefully"
// Change message
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

/* This is me testing godocs */
func LeTestMap () interface{}{
	something := map[string]string{"name": "Test"}
	return something
}

/* This is me just testing godocs too */
func LeTestStruct () interface{}{
	type nested struct {
		name string;
		age int;
	}

	type testt struct {
		name string;
		age int;
		nes nested;
	}
	
	
	something := testt{"a name", 15, nested{"hmm", 20}}
	return something;
}