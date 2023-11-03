package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPokemons(c echo.Context) error {
	return c.Render(http.StatusOK, "pokemons/index", map[string]interface{}{
		"pageTitle":   "Echo CRUD | ポケモン一覧",
		"currentPath": c.Path(),
	})
}
