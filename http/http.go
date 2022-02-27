package http

import (
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Start() {
	r.Run()
}
