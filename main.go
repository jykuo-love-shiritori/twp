package main

import (
	"github.com/jykuo-love-shiritori/twp/pkg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Group("/", middleware.Static("frontend/dist"))
	pkg.RouterHandler(e.Group("/api"))

	e.Logger.Fatal(e.Start(":8080"))
}
