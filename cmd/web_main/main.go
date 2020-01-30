package main

import (
	"github.com/labstack/echo/v4"
	"go.guoyk.net/env"
	"nautilus/pkg/exe"
	"nautilus/pkg/foxtrot"
	"net/http"
)

var (
	optBind string
)

func setup() (err error) {
	if err = env.StringVar(&optBind, "BIND", ":4000"); err != nil {
		return
	}
	return
}

func main() {
	var err error
	defer exe.Exit(&err)

	exe.Project = "web_main"
	exe.Setup()

	if err = setup(); err != nil {
		return
	}

	f := foxtrot.New(foxtrot.Options{Addr: optBind})

	f.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "网站正在维护中\n京ICP备15056756号-1")
	})

	err = exe.RunFoxtrot(f)
}
