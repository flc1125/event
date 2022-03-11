package provider

import (
	"github.com/flc1125/event/app/http/controller"
)

func init() {
	api()
}

func api() {
	api := Engine.Group("/api")

	api.GET("/ping", controller.ApiController.Ping)
}
