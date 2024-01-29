package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userRouter := r.Group("/user")
	userRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"svc": "api",
		})
	})

	r.Run(":1001")
}
