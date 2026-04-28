package main

import (
	db "design-pattern/proxy/config/database"
	"design-pattern/proxy/model"
	"fmt"
)

func main() {
	db.Connect()

	// lazy profile
	user := model.NewUser(1)
	// ProxyProfile.load() -> Profile.Get -> Profile.load
	fmt.Println(user.Profile.Get())

}
