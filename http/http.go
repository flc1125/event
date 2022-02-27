package http

import (
	"github.com/gin-gonic/gin"
)

var engine = gin.Default()

func Start() {
	engine.Run()
}
