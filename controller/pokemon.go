package controller

import (
	"net/http"

	"github.com/kyoruni/echo-crud/model"
	"github.com/labstack/echo/v4"
)

func GetPokemons(c echo.Context) error {
	pokemons := []model.Pokemon{}
	db, err := model.GetDB()
	if err != nil {
		return nil
	}

	db.Find(&pokemons)
	return c.Render(http.StatusOK, "pokemons/index", map[string]interface{}{
		"pageTitle":   "Echo CRUD | ポケモン一覧",
		"currentPath": c.Path(),
		"pokemons":    pokemons,
	})
}
