package main

import "golang.org/x/crypto/bcrypt"

type user struct {
	username string
	password password
}

func newUsr(un, pw string) user {
	user := user{username: un, password: newPassword(pw)}
	return user
}

type password struct {
	password string
}

func newPassword(pw string) password {
	p := password{password: pw}
	return p
}

type saltedPwd struct {
	saltedPwd []byte
}

func (p password) saltPwd() saltedPwd {
	var spwd saltedPwd
	var err error
	spwd.saltedPwd, err = bcrypt.GenerateFromPassword([]byte(p.password), 14)
	if err != nil {
		panic("couldnt salt password: " + err.Error())
	}
	return spwd
}

func (sp saltedPwd) verifyPassword(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(sp.saltedPwd), []byte(p))
	if err != nil {
		return false
	} else {
		return true
	}
}
