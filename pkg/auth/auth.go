package auth

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(Username string, key string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": Username,
		},
	)
	token, err := t.SignedString([]byte(key))
	if err != nil {
		log.Println(err)
	}
	return token
}
