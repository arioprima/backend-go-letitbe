package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, map[string]string{
			"message":    "Hello World Dong!",
			"coba layer": "http://localhost:8081/layer",
		})
	})

	err := router.Run(":8081")
	if err != nil {
		return
	}
}
