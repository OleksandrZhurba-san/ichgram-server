package server

import (
	"os"

	"github.com/OleksandrZhurba-san/ichgram-server/internal/router"
)

func Init() {

	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)

}
