package model

import "testing"

func TestUser_AddUser(t *testing.T) {
	user := User{
		UserName: "wq",
		Password: "11",
		Email:    "23",
	}
	err := user.AddUser()
	if err != nil {
		panic(err)
	}
}
