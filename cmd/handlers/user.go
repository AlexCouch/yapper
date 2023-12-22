package handlers

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/auth"
	"yapper.com/m/yapper/cmd/models"
	"yapper.com/m/yapper/cmd/repositories"
)

type UserProfileData struct{
    User models.User
    Pfp string
    PfpLen int64
    SignedIn bool
}

func UserProfile(c echo.Context) error {
    id := c.Param("id")
    idInt, err := strconv.Atoi(id)
    
    userData := UserProfileData{
        SignedIn: false,
    }

    user, err := repositories.GetUserById(idInt)
    if err != nil{
        //Replace with html template displaying error
        return c.String(http.StatusOK, err.Error())
    }
    userData.User = user
    pfp, err := repositories.GetProfilePictureById(user.ProfileImage)
    if pfp.Id > 0{
        imgbase64 := base64.StdEncoding.EncodeToString(pfp.Bytes)
        userData.Pfp = imgbase64
        userData.PfpLen = pfp.Length
    }

    signedIn, err := auth.CheckSignedIn(idInt, c.Cookies())
    if err != nil{
        return err
    }
    if signedIn{
        userData.SignedIn = true
    }
    return c.Render(http.StatusOK, "profile", userData)
}
