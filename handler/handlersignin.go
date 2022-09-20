package handler

import (
	"mahjong_sns/infra/mysql"
	"mahjong_sns/utility"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Handlersignin(c *gin.Context) {
	user, err := mysql.UserGetOneByAccount(c.PostForm("account"))
	hashedPW := utility.HashStr(c.PostForm("password"))
	if err != nil {
		c.Redirect(303, "/signup?err=acc")
	} else if user.Password != hashedPW {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{"error": "Password is incorrect."})
	} else {
		var id string = strconv.FormatUint(uint64(user.ID), 10)
		url := "/" + id
		c.Redirect(302, url)
	}
}
