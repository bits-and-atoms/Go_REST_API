package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email string, userId int64) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == ""{
		return "",errors.New("secret signing key not found")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	rel, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return rel, nil
}
