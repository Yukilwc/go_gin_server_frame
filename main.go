package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("22", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "33"})
	})
	r.Run()
}
