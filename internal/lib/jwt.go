package lib

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret123")

func GenerateToken(user_id int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Minute * 10).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signin method tidak valid")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, errors.New("token tidak valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil

}
