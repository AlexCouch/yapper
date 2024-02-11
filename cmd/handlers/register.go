package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/models"
	"yapper.com/m/yapper/cmd/repositories"
	"yapper.com/m/yapper/cmd/views"
)

func GotoRegister(c echo.Context) error {
	comp := views.Register(views.RegisterInfo{})
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func Register(c echo.Context) error {
	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := repositories.CreateUser(models.User{
		Email:    email,
		Name:     username,
		Password: password,
	})
	if err != nil {
		comp := views.Register(views.RegisterInfo{})
		return comp.Render(c.Request().Context(), c.Response().Writer)
	}

	return c.Redirect(http.StatusMovedPermanently, "/user/"+fmt.Sprint(user.Id))
}
