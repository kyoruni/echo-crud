package route

import (
	"github.com/kyoruni/echo-crud/controller"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()
	e.GET("/", controller.GetPokemons)
	e.GET("/pokemons", controller.GetPokemons)
	e.GET("/types", controller.GetTypes)
	return e
}
