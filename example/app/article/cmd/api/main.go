package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userRouter := r.Group("/article")
	userRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"svc": "article",
		})
	})

	r.Run(":1002")
}
