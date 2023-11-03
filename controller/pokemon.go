package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPokemons(c echo.Context) error {
	return c.Render(http.StatusOK, "pokemons/index", nil)
}
