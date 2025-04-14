package user

import (
	"github.com/OleksandrZhurba-san/ichgram-server/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(r *gin.RouterGroup, db *mongo.Database) {
	repo := NewUserRepository(db)
	handle := NewHanlder(repo)

	r.POST("/register", handle.Register)
	r.POST("/login", handle.LoginUser)
	r.GET("/:id", middleware.ValidateObjectID("id"), handle.GetUserByID)
}
