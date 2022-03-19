package token

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateJWT(user_id string) (string, error) {
	token_life, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFE"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_life)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
