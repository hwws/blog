package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        //继承Model
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
}
