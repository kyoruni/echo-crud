package route

import (
	"html/template"

	"github.com/kyoruni/echo-crud/controller"
	"github.com/kyoruni/echo-crud/utils"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()

	// テンプレートを登録
	t := template.Must(template.ParseGlob("views/*/*.html"))
	e.Renderer = &utils.Template{
		Templates: t,
	}

	e.GET("/", controller.GetPokemons)
	e.GET("/pokemons", controller.GetPokemons)
	e.GET("/types", controller.GetTypes)
	return e
}
