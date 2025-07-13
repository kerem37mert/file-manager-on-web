package main

import (
	"filemanager/handlers"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Static("/static", "static")

	renderer := &TemplateRenderer{template.Must(template.ParseGlob("templates/*.html"))}
	e.Renderer = renderer

	e.GET("/*", handlers.Page)
	e.GET("/filemanagerapi/new", handlers.New)

	e.Logger.Fatal(e.Start(":1881"))
}
