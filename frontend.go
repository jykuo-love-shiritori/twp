package main

import (
	"strings"

	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterFrontend(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().URL.Path, constants.SWAGGER_PATH) {
				return true
			}

			if strings.HasPrefix(c.Request().URL.Path, constants.API_BASE_PATH) {
				return true
			}

			return false
		},
		Root:  "frontend/dist",
		HTML5: true,
	}))
}
