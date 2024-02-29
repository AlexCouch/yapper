package models

import "time"

type User struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Bio          string    `json:"bio"`
	ProfileImage int       `json:"profile_image"`
	Banner       int       `json:"banner"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ProfileImage struct {
	Id     int    `json:"id"`
	Path   string `json:"path"`
	Length int64  `json:"len"`
	Bytes  []byte `json:"bytes"`
}

type Yap struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"create_at"`
}
