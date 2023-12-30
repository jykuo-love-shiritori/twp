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
			skipList := []string{constants.SWAGGER_PATH, constants.API_BASE_PATH, constants.IMAGE_BASE_PATH}

			for _, v := range skipList {
				if strings.HasPrefix(c.Request().URL.Path, v) {
					return true
				}
			}

			return false
		},
		Root:  "frontend/dist",
		HTML5: true,
	}))
}
