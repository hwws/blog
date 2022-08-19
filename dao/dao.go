package dao

import (
	"blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	AddUser(user *model.User)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	// gorm连接数据库
	dsn := "root:123456@tcp(192.168.18.129:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //连接远程主机的folang_db数据库
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db} // 创建数据库, 数据库赋值给接口Mgr, 隐藏db, 确保只有一个数据库
	db.AutoMigrate(&model.User{})
}

func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user) //insert
}
