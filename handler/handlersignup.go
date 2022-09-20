package handler

import (
	"github.com/gin-gonic/gin"
)

func Handlersignup(c *gin.Context) {
	err := c.Query("err")
	if err == "acc" {
		err = "The account name is not exist, so you should make a new account."
	}
	c.HTML(200, "signup.html", gin.H{"error": err})
}
