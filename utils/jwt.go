package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email string, userId int64) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("secret signing key not found")
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

func VerifyToken(token string) (int64,error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if ok == false {
			return nil, errors.New("unexpected signin method")
		}
		secretKey := os.Getenv("SECRET_KEY")

		return []byte(secretKey), nil
	})
	if err != nil {
		return 0,errors.New("could not parse token")
	}
	isok := parsedToken.Valid
	if isok == false {
		return 0,errors.New("invalid token")
	}
	// _, ok := parsedToken.Claims.(jwt.MapClaims)
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0,errors.New("invalid token claims")
	}
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId,nil
}
