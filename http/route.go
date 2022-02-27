package http

import (
	"github.com/gin-gonic/gin"
)

func init() {
	api()
}

func api() {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
