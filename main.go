package main

import (
	"log"
	"web-native/model"
)

func main() {
	user := model.User{
		UserName: "KKKa1",
		Password: "0901",
		Email:    "1111",
	}
	err := user.AddUser()
	if err != nil {
		log.Println(err)
	}
}
