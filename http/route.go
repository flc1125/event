package http

import (
	"github.com/flc1125/event/http/controller"
)

func init() {
	api()
}

func api() {
	api := engine.Group("/api")

	apiController := &controller.ApiController{}

	api.GET("/ping", apiController.Ping)
}
