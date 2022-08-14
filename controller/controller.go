package controller

import (
	"blog/dao"
	"blog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)
}

func ListUser(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
