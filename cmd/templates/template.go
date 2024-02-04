package templates

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates map[string]*template.Template
}

func NewTemplate() *Template {
	return &Template{
		templates: map[string]*template.Template{},
	}
}

func (t *Template) normalizePaths(paths []string) []string {
	normals := []string{}
	for _, path := range paths {
		normal := filepath.Join("public/views/", path+".html")
		normals = append(normals, normal)
	}
	return normals
}

func (t *Template) AddTemplate(name string, children ...string) {
	path := filepath.Join("public/views/", name+".html")
	children = t.normalizePaths(children)
	paths := []string{}
	paths = append(paths, path)
	paths = append(paths, children...)
	paths = append(paths, "public/views/base.html")
	t.templates[name] = template.Must(template.ParseFiles(paths...))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	templ, ok := t.templates[name]
	if !ok {
		return errors.New("Template " + name + " does not exist in template map!")
	}
	err := templ.ExecuteTemplate(w, "base", data)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return nil
}
