package repositories

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"yapper.com/m/yapper/cmd/models"
	"yapper.com/m/yapper/cmd/storage"
	//"github.com/go-pg/pg/v10/orm"
)

func CreateUser(user models.User) (models.User, error) {
	db, err := storage.GetDB()
	if err != nil {
		return user, err
	}
	now := time.Now()
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(pass)
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err = db.Model(&user).Insert()
	if err != nil {
		return user, err
	}
	return user, nil
}

type incorrectPassword struct{}

var IncorrectPassword = incorrectPassword{}

func (i incorrectPassword) Error() string {
	return "Incorrect password!"
}

type incorrectEmail struct{}

var IncorrectEmail = incorrectEmail{}

func (i incorrectEmail) Error() string {
	return "Incorrect email!"
}

func GetUserById(id int) (models.User, error) {
	db, err := storage.GetDB()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	user := models.User{Id: id}
	err = db.Model(&user).WherePK().Select()

	return user, err
}

func GetProfilePictureById(id int) (models.ProfileImage, error) {
	db, err := storage.GetDB()
	if err != nil {
		return models.ProfileImage{}, err
	}
	defer db.Close()

	pfp := models.ProfileImage{Id: id}
	err = db.Model(&pfp).WherePK().Select()

	return pfp, err
}

func GetUser(email string, password string) (models.User, error) {
	db, err := storage.GetDB()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	user := models.User{}
	err = db.Model(&user).
		Where("email = ?", email).
		Select()
	if err != nil {
		return user, IncorrectEmail
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, IncorrectPassword
	}
	return user, err
}

func DeleteUser(id int) error {
	db, err := storage.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Model(&models.User{Id: id}).WherePK().Delete()
	return err
}

func UploadProfilePicture(pfp models.ProfileImage) (int, error) {
	db, err := storage.GetDB()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	var id int
	_, err = db.Model(&pfp).Returning("id").Insert(&id)
	println(id)

	return id, err
}

func UpdateUser(user models.User, column string) error {
	db, err := storage.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Model(&user).WherePK().Column(column).Update()
	return err
}
