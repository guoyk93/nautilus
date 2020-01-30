package foxtrot

import (
	"context"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/http/pprof"
)

type Options struct {
	Addr string
}

type Foxtrot struct {
	*echo.Echo
	Addr string
}

func New(opts Options) *Foxtrot {
	if len(opts.Addr) == 0 {
		opts.Addr = ":4000"
	}
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	// pprof
	e.Any("/debug/pprof/", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	e.Any("/debug/pprof/cmdline", echo.WrapHandler(http.HandlerFunc(pprof.Cmdline)))
	e.Any("/debug/pprof/profile", echo.WrapHandler(http.HandlerFunc(pprof.Profile)))
	e.Any("/debug/pprof/symbol", echo.WrapHandler(http.HandlerFunc(pprof.Symbol)))
	e.Any("/debug/pprof/trace", echo.WrapHandler(http.HandlerFunc(pprof.Trace)))
	// log
	e.Use(NewLogger())
	// recover
	e.Use(middleware.Recover())
	// metrics
	p := prometheus.NewPrometheus("http", nil)
	p.Use(e)
	return &Foxtrot{
		Echo: e,
		Addr: opts.Addr,
	}
}

func (f *Foxtrot) Start(ech chan error) {
	go func() {
		if ech == nil {
			_ = f.Echo.Start(f.Addr)
		} else {
			ech <- f.Echo.Start(f.Addr)
		}
	}()
}

func (f *Foxtrot) Shutdown(ctx context.Context) error {
	return f.Echo.Shutdown(ctx)
}
