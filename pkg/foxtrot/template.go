package foxtrot

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"path/filepath"
)

type Template struct {
	dir       string
	reload    bool
	templates *template.Template
}

func NewTemplate(dir string, reload bool) *Template {
	return &Template{
		dir:       dir,
		reload:    reload,
		templates: template.Must(template.ParseGlob(filepath.Join(dir, "*.html"))),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if t.reload {
		t.templates = template.Must(template.ParseGlob(filepath.Join(t.dir, "*.html")))
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
