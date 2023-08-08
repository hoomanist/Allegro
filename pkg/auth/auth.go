package auth

import (
	"errors"
	"log"
	"net/http"

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

func Authorize(endpointHandler func(w http.ResponseWriter, r *http.Request), key string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				log.Println(r.Header["Token"][0])
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					log.Println(1)
					http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
				}
				return []byte(key), nil

			})
			if err != nil {
				log.Println(err)
				http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
				return
			}
			if token.Valid {
				endpointHandler(w, r)
			} else {
				log.Println(3)
				http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
				return
			}
		} else {
			log.Println(4)
			http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
			return
		}
	})
}

func ExtractClaims(_ http.ResponseWriter, r *http.Request, key string) (string, error) {
	if r.Header["Token"] != nil {
		tokenString := r.Header["Token"][0]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("there's an error with the signing method")
			}
			return []byte(key), nil
		})
		if err != nil {
			return "", err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			username := claims["user"].(string)
			return username, nil
		}
	}

	return "", errors.New("unable to extract claims")
}
