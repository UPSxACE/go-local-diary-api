package main

import (
	"html/template"

	"github.com/UPSxACE/go-local-diary-api/controllers"
	"github.com/UPSxACE/go-local-diary-api/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Pre-compile templates in views subdirectories, and subdirectories of those subdirectories
	tBuilder :=  template.Must(template.ParseGlob("views/*/*.html"));
	tBuilder = template.Must(tBuilder.ParseGlob("views/*/*/*.html"))
	t := &views.Template{
		Templates: tBuilder,
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files
	e.Static("/public", "public")

	e.Renderer = t

	// Routes
	e.GET("/", controllers.IndexController)
	e.POST("/htmx/change-message", controllers.HtmxChangeMessageController)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
