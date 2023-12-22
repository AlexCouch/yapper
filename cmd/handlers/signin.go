package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/auth"
	"yapper.com/m/yapper/cmd/repositories"
)

type ReturnData struct{
    EmailSuccess bool
    PasswordSuccess bool
}

func GotoSignin(c echo.Context) error {
    return c.Render(http.StatusOK, "signin", ReturnData{EmailSuccess: true, PasswordSuccess: true})
}

func Signin(c echo.Context) error {
    email := c.FormValue("email")
    password := c.FormValue("password")
    user, err := repositories.GetUser(email, password)
    if err != nil{
        returnData := ReturnData{
            EmailSuccess: true,
            PasswordSuccess: true,
        }
        if errors.Is(err, repositories.IncorrectEmail){
            returnData.EmailSuccess = false
        }else if errors.Is(err, repositories.IncorrectPassword){
            returnData.PasswordSuccess = false
        }else{
            return c.String(http.StatusOK, "Unknown error: " + err.Error())
        }
        return c.Render(http.StatusOK, "signin", returnData)
    }
    cookie := new(http.Cookie)
    cookie.Name = "user"
    token, err := auth.CreateJWT(user.Id)
    if err != nil{
        return err
    }
    cookie.Value = token
    c.SetCookie(cookie)
    return c.Redirect(http.StatusMovedPermanently, "user/" + fmt.Sprint(user.Id))
}
