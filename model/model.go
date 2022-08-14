package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        //继承Model
	Username   string `json:"username"`
	Password   string `json:"password"`
}
