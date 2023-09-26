package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RouterHandler(e *echo.Group) {
	e.GET("", func(c echo.Context) error { return c.String(http.StatusOK, "ok") })
	e.GET("/hello", func(c echo.Context) error { return c.JSON(http.StatusOK, map[string]any{"message": "ok"}) })
}
