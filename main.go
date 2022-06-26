package main

import (
	controller "golang-app/app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthcheck", controller.HealthCheck)
	r.GET("/hello", controller.Hello)
	r.Run(":8080")
}
