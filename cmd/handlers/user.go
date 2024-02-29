package handlers

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/auth"
	"yapper.com/m/yapper/cmd/repositories"
	"yapper.com/m/yapper/cmd/views"
)

func UserProfile(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	userData := views.ProfileInfo{
		SignedIn: false,
	}

	user, err := repositories.GetUserById(idInt)
	if err != nil {
		//Replace with html template displaying error
		return c.String(http.StatusOK, err.Error())
	}
	userData.User = user
	pfp, err := repositories.GetProfilePictureById(user.ProfileImage)
	if pfp.Id > 0 {
		imgbase64 := base64.StdEncoding.EncodeToString(pfp.Bytes)
		userData.Pfp = imgbase64
		userData.PfpLen = pfp.Length
	}

	signedIn, err := auth.CheckUserIdSignedIn(idInt, c.Cookies())
	if err != nil {
		return err
	}
	if signedIn {
		userData.SignedIn = true
	}
	comp := views.Profile(userData)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}
