package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authorize(c echo.Context) error {
	params := make(map[string]any)

	c.Bind(&params)

	b, _ := json.MarshalIndent(params, "", "  ")
	fmt.Println(string(b))

	return c.NoContent(http.StatusOK)
}
