package handlers

import (
	"errors"

	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/views"
)

type YapListData struct {
	Count int
}

func YapList(ctx echo.Context) error {
	if !htmx.IsHTMX(ctx.Request()) {
		return errors.New("Expected an HTMX request")
	}
	return htmx.NewResponse().
		RenderTempl(ctx.Request().Context(), ctx.Response().Writer, views.Yaplist())
}
