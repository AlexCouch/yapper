package handlers

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/angelofallars/htmx-go"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"yapper.com/m/yapper/cmd/auth"
	"yapper.com/m/yapper/cmd/models"
	"yapper.com/m/yapper/cmd/repositories"
	"yapper.com/m/yapper/cmd/views"
)

func EditProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.CustomClaim)
	user_id := claims.Id

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		//TODO: Make error page
		return err
	}
	if user_id != id {
		return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/user/%d", id))
	}
	target_user, err := repositories.GetUserById(id)
	if err != nil {
		return err
	}
	editData := views.EditProfileData{User: target_user}
	if target_user.ProfileImage > 0 {
		pfp, err := repositories.GetProfilePictureById(target_user.ProfileImage)
		if err != nil {
			return err
		}
		imgBase64 := base64.StdEncoding.EncodeToString(pfp.Bytes)
		editData.Pfp = imgBase64
		editData.PfpLen = pfp.Length
	}
	comp := views.EditProfile(editData)
	return htmx.NewResponse().
		RenderTempl(c.Request().Context(), c.Response().Writer, comp)
}

func uploadProfilePicture(pfp_header *multipart.FileHeader) (int, error) {
	bytes := make([]byte, pfp_header.Size)
	if pfp_header.Size > 0 {
		file, err := pfp_header.Open()
		if err != nil {
			return -1, err
		}
		reader := bufio.NewReader(file)
		_, err = reader.Read(bytes)
		if err != nil {
			return -1, err
		}
	}
	pfp := models.ProfileImage{
		Path:   pfp_header.Filename,
		Length: pfp_header.Size,
		Bytes:  bytes,
	}
	return repositories.UploadProfilePicture(pfp)
}

func SaveProfileChanges(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	name := c.FormValue("name")
	bio := c.FormValue("bio")
	pfp_header, err := c.FormFile("pimg")
	var pfp_id int
	if err == nil {
		pfp_id, err = uploadProfilePicture(pfp_header)
		if err != nil {
			return err
		}
	}
	user := models.User{Id: id, Name: name, Bio: bio, ProfileImage: pfp_id}
	err = repositories.UpdateUser(user, "name")
	if err != nil {
		return err
	}
	err = repositories.UpdateUser(user, "bio")
	if err != nil {
		return err
	}
	if pfp_id > 0 {
		err = repositories.UpdateUser(user, "profile_image")
		if err != nil {
			return err
		}
	}
	return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/user/%d", id))
}

func uploadBanner(header *multipart.FileHeader) (models.Banner, error) {
	size := header.Size
	data := make([]byte, size)
	if size == 0 {
		return models.Banner{}, nil
	}
	file, err := header.Open()
	if err != nil {
		return models.Banner{}, err
	}
	defer file.Close()

	bufReader := bufio.NewReader(file)
	_, err = bufReader.Read(data)
	if err != nil {
		return models.Banner{}, err
	}
	banner := models.Banner{
		Data:   data,
		Length: int(size),
	}
	bannerId, err := repositories.UploadBanner(banner)
	if err != nil {
		return models.Banner{}, err
	}
	banner.Id = bannerId
	return banner, nil
}

func EditBanner(c echo.Context) error {
	if !htmx.IsHTMX(c.Request()) {
		//Change to an error dialog/modal/toast instead
		return errors.New("Expected an htmx request")
	}
	userCookie, err := c.Cookie("user")
	if err != nil {
		println(err)
		return err
	}
	if err := auth.CheckSignedIn(userCookie); err != nil {
		println(err)
		return err
	}
	userId, err := auth.GetUser(userCookie.Value)
	if err != nil {
		return err
	}

	header, err := c.FormFile("banner")
	if err != nil {
		return err
	}
	banner, err := uploadBanner(header)
	if err != nil {
		return err
	}

	user := models.User{Id: userId, Banner: banner.Id}
	err = repositories.UpdateUser(user, "banner")
	if err != nil {
		return nil
	}
	bannerB64 := base64.StdEncoding.EncodeToString(banner.Data)

	comp := views.Banner(bannerB64)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, comp)
}
