package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/auth"
	"yapper.com/m/yapper/cmd/repositories"
	"yapper.com/m/yapper/cmd/views"
)

func GotoSignin(c echo.Context) error {
	info := views.LoginInfo{
		EmailSuccess:    true,
		PasswordSuccess: true,
	}
	comp := views.Login(info)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := repositories.GetUser(email, password)
	if err != nil {
		info := views.LoginInfo{
			EmailSuccess:    true,
			PasswordSuccess: true,
		}
		if errors.Is(err, repositories.IncorrectEmail) {
			info.EmailSuccess = false
		} else if errors.Is(err, repositories.IncorrectPassword) {
			info.PasswordSuccess = false
		} else {
			return c.String(http.StatusOK, "Unknown error: "+err.Error())
		}
		comp := views.Login(info)
		return comp.Render(c.Request().Context(), c.Response().Writer)
	}
	cookie := new(http.Cookie)
	cookie.Name = "user"
	token, err := auth.CreateJWT(user.Id)
	if err != nil {
		return err
	}
	cookie.Value = token
	c.SetCookie(cookie)
	return c.Redirect(http.StatusMovedPermanently, "user/"+fmt.Sprint(user.Id))
}
