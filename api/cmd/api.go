package main

import (
	"log"

	"github.com/Robert076/logger2.git/api/pkg
	/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Hello World from CI
	r := gin.New()
	r.GET("/", handlers.HandlerGet)
	r.POST("/", handlers.HandlerPost)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start http server")
	}
}
