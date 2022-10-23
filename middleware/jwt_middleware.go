package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId int, email string) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
