package server

import (
	"os"

	"github.com/OleksandrZhurba-san/ichgram-server/common/db"
	"github.com/OleksandrZhurba-san/ichgram-server/internal/router"
	"github.com/gin-gonic/gin"
)

func Init() {
	client := db.Init()
	database := client.Database("ichgram")

	r := gin.Default()
	api := r.Group("/api")

	router.InitRoutes(api, database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)

}
