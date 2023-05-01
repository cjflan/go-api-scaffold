package main

import (
	"github.com/cjflan/go-api-scaffold/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	add_routes(router)

	router.Run(":8080")
}

func add_routes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
}
