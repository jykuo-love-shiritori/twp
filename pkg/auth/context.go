package auth

import (
	"github.com/labstack/echo/v4"
)

const authContextKey string = "key"

// To use this function the endpoint must have
// used the `IsRole` or `ValidateJwt` middleware
func GetUsername(c echo.Context) (string, bool) {
	username, ok := c.Get(authContextKey).(string)
	return username, ok
}
