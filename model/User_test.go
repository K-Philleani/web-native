package model

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试开始")
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("测试User相关方法")
	//t.Run("测试添加用户",TestUser_AddUser)
	t.Run("测试单条查询", TestUser_GetUserById)
	t.Run("测试多条查询", TestUser_GetUsers)
}

func TestUser_AddUser(t *testing.T) {
	user := User{
		UserName: "w1q",
		Password: "111",
		Email:    "213",
	}
	err := user.AddUser()
	if err != nil {
		panic(err)
	}
}

func TestUser_GetUserById(t *testing.T) {
	user := User{Id: 3}
	userInfo, err := user.GetUserById()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", userInfo)
}

func TestUser_GetUsers(t *testing.T) {
	user := User{Id: 0}
	userList, err := user.GetUsers()
	if err != nil {
		panic(err)
	}
	for k, v := range userList {
		fmt.Printf("第%d条, %v\n", k+1, v)

	}
}


