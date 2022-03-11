package provider

import "github.com/gin-gonic/gin"

var Engine = gin.Default()

func Run() {
	Engine.Run()
}
