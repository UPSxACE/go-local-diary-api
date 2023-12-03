package main

import (
	"html/template"

	"github.com/UPSxACE/go-local-diary-api/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Pre-compile templates
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = t

	// Routes
	e.GET("/", api.Index)
	e.POST("/change-message", api.ChangeMessage)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
