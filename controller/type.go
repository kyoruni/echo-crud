package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTypes(c echo.Context) error {
	return c.String(http.StatusOK, "get types")
}
