package controller

import (
	"github.com/gin-gonic/gin"
)

var ApiController = &apiController{}

type apiController struct{}

func (this *apiController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
