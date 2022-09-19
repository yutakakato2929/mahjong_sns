package handler

import (
	"mahjong_sns/infra/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handleruserinsert(c *gin.Context) {
	var form mysql.User
	if err := c.Bind(&form); err != nil {
		users := mysql.UserGetAll()
		c.HTML(http.StatusBadRequest, "index.html", gin.H{"users": users, "err": err})
		c.Abort()
	} else {
		acc := c.PostForm("account")
		name := c.PostForm("name")
		pw := c.PostForm("password")
		mysql.UserInsert(acc, name, pw)
		c.Redirect(302, "/")
	}
}
