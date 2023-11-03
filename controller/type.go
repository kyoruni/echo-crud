package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTypes(c echo.Context) error {
	return c.Render(http.StatusOK, "types/index", map[string]interface{}{
		"pageTitle": "Echo CRUD | タイプ一覧",
	})
}
