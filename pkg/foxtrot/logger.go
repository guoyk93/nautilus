package foxtrot

import "github.com/labstack/echo/v4"

func NewLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			//TODO: integrate with zerolog
			return next(c)
		}
	}
}
