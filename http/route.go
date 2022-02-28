package http

import (
	"github.com/flc1125/event/http/controller"
)

func init() {
	api()
}

func api() {
	api := engine.Group("/api")

	api.GET("/ping", controller.ApiController.Ping)
}
