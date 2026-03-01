package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
	server := gin.Default()
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{"message":"healthy"})
	})
	server.Run(":8080")
}