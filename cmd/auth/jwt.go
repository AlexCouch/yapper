package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaim struct{
    Id int
    jwt.RegisteredClaims
}

func CreateJWT(user int) (string, error){
    claims := CustomClaim{
        user,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte("secret"))
    if err != nil{
        return "", err
    }
    return t, nil
}

