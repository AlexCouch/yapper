package templates

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Template struct{
    templates map[string]*template.Template
}

func NewTemplate() *Template{
    return &Template{
        templates: map[string]*template.Template{},
    }
}

func (t *Template) AddTemplate(name string) {
    path := filepath.Join("public/views/", name + ".html")
    t.templates[name] = template.Must(template.ParseFiles("public/views/base.html", "public/views/topbar.html", path))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
    templ, ok := t.templates[name]
    if !ok{
        return errors.New("Template " + name + " does not exist in template map!")
    }
    err := templ.ExecuteTemplate(w, "base", data)
    if err != nil{
        return c.String(http.StatusOK, err.Error())
    }
    return nil
}
