package main

import "fmt"

func main() {
	usr := newUsr("things", "wafaw")
	fmt.Println(usr.username)
	tkns := createToken(usr)
	fmt.Println(tkns)
	tkn, _ := verifyToken(tkns)

}
