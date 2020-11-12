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

var Db = utils.Db

func (user *User) AddUser() error{
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


func (user *User) GetUserById() (userInfo User, err error) {
	sqlStr := `select id, username, password, email from Users where id = ?`
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		return
	}
	row := stmt.QueryRow(user.Id)
	var id int64
	var username string
	var password string
	var email string
	err = row.Scan(&id,& username, &password,&email)
	if err != nil {
		return
	}
	userInfo.Id = id
	userInfo.UserName = username
	userInfo.Password = password
	userInfo.Email = email
	return
}

func (user *User) GetUsers() (userList []User, err error) {
	sqlStr := `select id, username, password, email from Users where id > ?`
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		return
	}
	rows, err := stmt.Query(user.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var username string
		var password string
		var email string
		err = rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return
		}
		u := User{
			Id:       id,
			UserName: username,
			Password: password,
			Email:    email,
		}
		userList = append(userList, u)
	}
	return
}