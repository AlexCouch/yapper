package auth

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func CheckSignedIn(userId int, cookies []*http.Cookie) (bool, error){
    for cid := range cookies{
        cookie := cookies[cid]
        //Replace with JWT
        if cookie.Name == "user"{
            token := cookie.Value
            tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
                if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
                    return nil, fmt.Errorf("Signing method incorrect")
                }
                return []byte("secret"), nil
            })
            if err != nil{
                return false, err
            }
            claim, ok := tkn.Claims.(jwt.MapClaims)
            if !ok{
                return false, fmt.Errorf("Unable to get MapClaims from token. Report to developer")
            }
            if !tkn.Valid{
                return false, fmt.Errorf("Token invalid")
            }
            id, ok := claim["Id"].(float64)
            if !ok {
                return false, fmt.Errorf("Unable to get Id from token header")
            }
            return userId == int(id), nil
        }
    }
    return false, nil
}
