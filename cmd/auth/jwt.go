package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaim struct {
	Id int
	jwt.RegisteredClaims
}

func CreateJWT(user int) (string, error) {
	claims := CustomClaim{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetUser(token string) (int, error) {
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method incorrect")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return -1, err
	}
	if !tkn.Valid {
		return -1, fmt.Errorf("Token invalid")
	}
	claim, ok := tkn.Claims.(jwt.MapClaims)
	if !ok {
		return -1, fmt.Errorf("Unable to get MapClaims from token. Report to developer")
	}

	id, ok := claim["Id"].(float64)
	if !ok {
		return -1, fmt.Errorf("Unable to get Id from token header")
	}
	return int(id), nil
}
