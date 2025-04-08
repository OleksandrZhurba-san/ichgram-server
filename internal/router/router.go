package router

import (
	"github.com/OleksandrZhurba-san/ichgram-server/internal/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(r *gin.RouterGroup, db *mongo.Database) {

	// User module
	userGroup := r.Group("/users")
	user.InitRoutes(userGroup, db)

}
