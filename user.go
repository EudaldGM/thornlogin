package main

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type user struct {
	username  string
	password  string
	saltedPwd string
}

func newUsr(un, pw string) user {
	user := user{username: un, password: pw}
	slog.Debug("New user created: " + un)
	return user
}

func (u *user) saltPwd() {
	spw, err := bcrypt.GenerateFromPassword([]byte(u.password), 14)
	u.password = ""
	u.saltedPwd = string(spw)
	if err != nil {
		panic("couldnt salt password: " + err.Error())
	}
	slog.Debug("Salted password of user: " + u.username)
}

func (u *user) verifyPassword(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.saltedPwd), []byte(p))
	if err != nil {
		slog.Error("User ", u.username, " verification failed")
		return false
	} else {
		slog.Info("User ", u.username, " verification successful")
		return true
	}
}
