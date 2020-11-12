package model

import (
	"log"
	"web-native/utils"
)

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}


func (user *User) AddUser() error{
	Db := utils.Db
	sqlStr := `insert into Users(username, password, email) values(?, ?, ?)`
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(user.UserName, user.Password, user.Email)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	log.Println("插入User的ID:", id)
	return nil
}