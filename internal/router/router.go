package router

import (
	"github.com/OleksandrZhurba-san/ichgram-server/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	user.RegisterRoutes(api)

	return r
}
