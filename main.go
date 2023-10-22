package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, map[string]string{
			"message": "Hello World!",
		})
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
