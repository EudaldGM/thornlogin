package main

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "secret"

func createToken(username user) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username.username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return []byte(secretKey), nil })
	if err != nil {
		return nil, err
	}
	return token, nil
}
