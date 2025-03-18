package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

const secretKey = "secret"

func createToken(u user) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.username,
		//get roles and assign them to the token
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	slog.Debug(fmt.Sprintf("Created token for user: %s", u.username))
	return tokenString
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return []byte(secretKey), nil })
	if err != nil {
		return nil, err
	}
	slog.Debug(fmt.Sprintf("Verified user %s token", token.Claims))
	return token, nil
}
