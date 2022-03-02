package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var ApiController = &apiController{}

type apiController struct{}

func (this *apiController) Ping(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "pong",
	//})
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Pong",
	})
}
