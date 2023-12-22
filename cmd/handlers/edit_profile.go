package handlers

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/auth"
	"yapper.com/m/yapper/cmd/models"
	"yapper.com/m/yapper/cmd/repositories"
)

type editProfileData struct{
    User models.User
    Pfp string
    PfpLen int64
}

func EditProfile(c echo.Context) error{
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*auth.CustomClaim)
    user_id := claims.Id

    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil{
        //TODO: Make error page
        return err
    }
    if user_id != id{
        return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/user/%d", id))
    }
    target_user, err := repositories.GetUserById(id)
    if err != nil{
        return err
    }
    editData := editProfileData{ User: target_user }
    if target_user.ProfileImage > 0{
        pfp, err := repositories.GetProfilePictureById(target_user.ProfileImage)
        if err != nil{
            return err
        }
        imgBase64 := base64.StdEncoding.EncodeToString(pfp.Bytes)
        editData.Pfp = imgBase64
        editData.PfpLen = pfp.Length
    }
    return c.Render(http.StatusOK, "edit_profile", editData)
}

func uploadProfilePicture(pfp_header *multipart.FileHeader) (int, error){
    bytes := make([]byte, pfp_header.Size)
    if pfp_header.Size > 0{
        file, err := pfp_header.Open()
        if err != nil{
            return -1, err
        }
        reader := bufio.NewReader(file)
        _, err = reader.Read(bytes)
        if err != nil{
            return -1, err
        }
    }
    pfp := models.ProfileImage{
        Path: pfp_header.Filename,
        Length: pfp_header.Size,
        Bytes: bytes,
    }
    return repositories.UploadProfilePicture(pfp)
}

func SaveProfileChanges(c echo.Context) error{
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil{
        return err
    }
    name := c.FormValue("name")
    bio := c.FormValue("bio")
    pfp_header, err := c.FormFile("pimg")
    var pfp_id int
    if err == nil{
        pfp_id, err = uploadProfilePicture(pfp_header)
        if err != nil{
            return err
        }
    }
    user := models.User{ Id: id, Name: name, Bio: bio, ProfileImage: pfp_id }
    err = repositories.UpdateUser(user, "name")
    if err != nil{
        return err
    }
    err = repositories.UpdateUser(user, "bio")
    if err != nil{
        return err
    }
    if pfp_id > 0{
        err = repositories.UpdateUser(user, "profile_image")
        if err != nil{
            return err
        }
    }
    return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/user/%d", id))
}
