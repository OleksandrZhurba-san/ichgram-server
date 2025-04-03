package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	userRoutes := rg.Group("/users")
	userRoutes.POST("/", CreateUser)
}
