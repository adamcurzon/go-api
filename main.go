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
			"message":   "Hello World",
			"time":      time.Now().Format(time.RFC850),
			"host":      "https://go.adamcurzon.co.uk",
			"framework": "gin",
		})
	})

	r.Run(":9990")
}
