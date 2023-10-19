package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func adminGetUser(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func adminDeleteUser(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func adminGetCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func adminAddCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func adminEditCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func adminDeleteCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func adminGetReport(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
