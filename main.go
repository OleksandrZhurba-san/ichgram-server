package main

import (
	"log"

	"github.com/OleksandrZhurba-san/ichgram-server/server"
	"github.com/joho/godotenv"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load", err)
	}

	server.Init()
}
