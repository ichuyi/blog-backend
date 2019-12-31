package util

import (
	"blog-backend/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyCustomClaims{
		UserId:   user.Id,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + MaxAge,
		},
	})
	return token.SignedString([]byte(Secret))
}
func ParseMyToken(t string) *MyCustomClaims {
	token, err := jwt.ParseWithClaims(t, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if token == nil || err != nil {
		return nil
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}
