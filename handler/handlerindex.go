package handler

import (
	"github.com/gin-gonic/gin"
)

func Handlerindex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
