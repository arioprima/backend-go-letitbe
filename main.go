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
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "Hello World",
			"coba layer": "http://localhost:8081/layer",
		})
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "Belajar CI/CD",
			"coba layer": "http://localhost:8081/belajar/cicd",
		})
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "Belajar NGINX/CD",
			"coba layer": "http://localhost:8081/belajar/cicd",
		})
	})

	err := router.Run(":8081")
	if err != nil {
		return
	}
}
