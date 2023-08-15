package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/pokemons", func(c echo.Context) error {
		return c.String(http.StatusOK, "pokemons index")
	})
	e.GET("/types", func(c echo.Context) error {
		return c.String(http.StatusOK, "types index")
	})
	return e
}
