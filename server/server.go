package server

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine{
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "welcome to my new home"})
	})

	return r
}
