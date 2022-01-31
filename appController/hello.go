package appController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "halo dunia 🌏")
}
