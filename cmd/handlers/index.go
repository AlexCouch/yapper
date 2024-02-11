package handlers

import (
	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/views"
)

func Home(c echo.Context) error {
	base := views.Base("Yapper")
	return base.Render(c.Request().Context(), c.Response().Writer)
}
