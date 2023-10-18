package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterFrontend(e *echo.Echo) {
	e.Group("/", middleware.Static("frontend/dist"))
}
