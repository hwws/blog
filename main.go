package main

import (
	"blog/dao"
	"blog/model"
)

func main() {
	user := model.User{
		Username: "hwws",
		Password: "123456",
	}

	dao.Mgr.AddUser(&user)
}
