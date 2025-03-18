package main

import "fmt"

func main() {
	usr := newUsr("things", "wafaw")
	usr.saltPwd()
	fmt.Println(usr)
	fmt.Println(usr.verifyPassword("wafaaw"))

}
