package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "testsecretkey"

type CustomClaim struct {
	Email  string
	UserID int64
}

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(token string) (*CustomClaim, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, errors.New("Unable to parse token")
	}

	if !parsedToken.Valid {
		return nil, errors.New("Invalid Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return &CustomClaim{
		Email:  email,
		UserID: userId,
	}, nil
}
