package handler

import (
	"github.com/gin-gonic/gin"
)

func Handlersignup(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}
