package handler

import (
	"mahjong_sns/infra/mysql"

	"github.com/gin-gonic/gin"
)

func Handlerindex(c *gin.Context) {
	users := mysql.UserGetAll()
	c.HTML(200, "index.html", gin.H{"users": users})
}
