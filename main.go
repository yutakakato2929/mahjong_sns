package main

import (
	"mahjong_sns/handler"
	"mahjong_sns/infra/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	mysql.DbInit()

	router.GET("/", handler.Handlerindex)
	router.POST("/signin", handler.Handlersignin)
	router.GET("/signup", handler.Handlersignup)
	router.POST("/userinsert", handler.Handleruserinsert)
	router.Run()
}
