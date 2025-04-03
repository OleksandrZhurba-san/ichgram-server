package main

import (
	"os"

	"github.com/OleksandrZhurba-san/ichgram-server/common/db"
	"github.com/OleksandrZhurba-san/ichgram-server/server"
)

func main() {

	db.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := server.Init()

	r.Run(":" + port)
}