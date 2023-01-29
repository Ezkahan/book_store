package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userId uint64) (string, error) {
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

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}

	return nil
}
