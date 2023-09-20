package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lhoffjann/goTodoApp/pkg/view"
)

type TemplateRegistry struct {
	templates *template.Template
}

func HomeHandler(c echo.Context) error {
	// Please note the the second parameter "home.html" is the template name and should
	// be equal to the value stated in the {{ define }} statement in "view/home.html"
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"name": "some things",
	})
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	tmpl := template.New("index")
	tmpl, _ = tmpl.Parse(view.Index)
	tmpl, _ = tmpl.Parse(view.Title)
	e.Renderer = &TemplateRegistry{
		templates: tmpl,
	}

	// Route => handler
	e.GET("/", HomeHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
