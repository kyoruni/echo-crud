package utils

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return template.Must(template.Must(t.Templates.Lookup("layout").Clone()).AddParseTree("content", t.Templates.Lookup(name).Tree)).ExecuteTemplate(w, "layout", data)
}
