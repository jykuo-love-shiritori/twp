package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Authorize(db *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		params := make(map[string]any)

		c.Bind(&params)

		b, _ := json.MarshalIndent(params, "", "  ")
		fmt.Println(string(b))

		return c.NoContent(http.StatusOK)
	}
}
