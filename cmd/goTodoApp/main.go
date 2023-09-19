package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lhoffjann/goTodoApp/pkg/view"
)

type Item struct {
	Name  string
	Count int
	Id    int
}
type Content struct {
	Items []Item
}
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	tmpl := template.New("index")
	var err error
	if tmpl, err = tmpl.Parse(view.Index); err != nil {
		fmt.Println(err)
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	items := Content{
		Items: []Item{
			{Name: "Item 1", Count: 1, Id: 1},
			{Name: "Item 2", Count: 2, Id: 2},
			{Name: "Item 3", Count: 3, Id: 3},
		},
	}
	e.Renderer = &TemplateRenderer{
		templates: tmpl,
	}
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", items)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
