package handler

import (
	"mahjong_sns/infra/mysql"
	"mahjong_sns/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handleruserinsert(c *gin.Context) {
	if c.PostForm("password") == c.PostForm("repassword") {
		_, err := mysql.UserGetOneByAccount(c.PostForm("account"))
		if err != nil {
			hashedPW := utility.HashStr(c.PostForm("password"))
			acc := c.PostForm("account")
			name := c.PostForm("name")
			mysql.UserInsert(acc, name, hashedPW)
			c.Redirect(302, "/")
		} else {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "The account name is used by an another user, so you should alter different account name."})
		}
	} else {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "You should retype the same password in the retyping box."})
	}
}

/*
if err := c.Bind(&form); err != nil {
		users := mysql.UserGetAll()
		c.HTML(http.StatusBadRequest, "index.html", gin.H{"users": users})
		c.Abort()
*/
