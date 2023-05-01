package main

import (
	"log"

	"github.com/cjflan/go-api-scaffold/controllers"

	"github.com/gin-gonic/gin"

	middleware "github.com/cjflan/go-api-scaffold/middleware"
)

func main() {
	router := gin.Default()

	add_routes(router)

	log.Print("Server listening on http://localhost:3000")
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}

}

func add_routes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
	router.GET("/validate", middleware.CheckJWTMiddle(), controllers.Login)
}
