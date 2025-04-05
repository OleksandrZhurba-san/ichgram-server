package main

import (
	"log"

	"github.com/OleksandrZhurba-san/ichgram-server/common/db"
	"github.com/OleksandrZhurba-san/ichgram-server/server"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load", err)
	}

	db.Init()

	server.Init()
}
