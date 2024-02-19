package handlers

import (
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/views"
)

func Home(c echo.Context) error {
	if htmx.IsHTMX(c.Request()) {
		comp := views.YaplistHTMX()
		return comp.Render(c.Request().Context(), c.Response().Writer)
	}
	base := views.Base("Yapper")
	return base.Render(c.Request().Context(), c.Response().Writer)
}
