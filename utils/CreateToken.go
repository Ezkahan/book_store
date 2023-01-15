package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId uint64) (string, error) {
	var err error
	claims := jwt.MapClaims{}
	claims["authorizer"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	newClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := newClaim.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
