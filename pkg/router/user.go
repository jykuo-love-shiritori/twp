package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func login(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func signup(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func logout(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func userGetInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func userEditInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func userUploadAvatar(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func userEditPassword(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func userGetCreditCard(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func userDeleteCreditCard(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func userAddCreditCard(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
