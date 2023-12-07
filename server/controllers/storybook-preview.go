package controllers

import "github.com/labstack/echo/v4"

func SetStorybookPreviewRoutes(e *echo.Echo) {
	e.GET("/storybook_preview/index-page-default", GetIndexController)
}

