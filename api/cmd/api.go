package main

import (
	"github.com/Robert076/logger2.git/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", handlers.HandlerGet)
	r.POST("/", handlers.HandlerPost)
	r.Run(":8080")
}
