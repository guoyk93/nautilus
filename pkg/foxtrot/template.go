package foxtrot

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"path/filepath"
)

type Template struct {
	templates *template.Template
}

func NewTemplate(dir string) *Template {
	return &Template{templates: template.Must(template.ParseGlob(filepath.Join(dir, "*.html")))}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
