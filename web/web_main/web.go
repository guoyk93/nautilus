package web_main

import (
	"github.com/labstack/echo/v4"
	"nautilus/svc/svc_id"
	"net/http"
)

type Web struct {
	ClientID *svc_id.Client
}

func (w *Web) Index(c echo.Context) (err error) {
	var ids []string
	if ids, err = w.ClientID.NextStr(c.Request().Context(), 1); err != nil {
		return
	}
	err = c.Render(http.StatusOK, "index", map[string]interface{}{
		"ID": ids[0],
	})
	return
}
