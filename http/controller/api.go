package controller

import (
	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (this *ApiController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
