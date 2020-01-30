package foxtrot

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Options struct {
	Addr string
}

type Foxtrot struct {
	*echo.Echo
}

func New(opts Options) *Foxtrot {
	e := echo.New()
	e.Use(NewLogger())
	e.Use(middleware.Recover())
	p := prometheus.NewPrometheus("http", nil)
	p.Use(e)
	return &Foxtrot{Echo: e}
}
