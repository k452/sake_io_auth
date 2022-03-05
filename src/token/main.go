package token

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateIDToken(user_id string) (string, error) {
	var hmacSampleSecret []byte
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user_id,
		"type": "id",
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

func CreateRefreshToken(user_id string) (string, error) {
	var hmacSampleSecret []byte
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user_id,
		"type": "refresh",
		"exp":  time.Now().AddDate(0, 0, 1).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

func ParseToken(tokenString string) interface{} {
	var hmacSampleSecret []byte
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return err
	}
}
