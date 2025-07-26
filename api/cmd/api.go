package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/", func(c *gin.Context) {
		// checks + write to db
	})
}
