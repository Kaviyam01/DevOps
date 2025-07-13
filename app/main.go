package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"ip":        c.ClientIP(),
		})
	})

	r.Run(":8080") // listen and serve on port 8080
}