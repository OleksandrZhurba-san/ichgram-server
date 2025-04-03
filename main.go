package main

import (
	"github.com/OleksandrZhurba-san/ichgram-server/common/db"
	"github.com/OleksandrZhurba-san/ichgram-server/server"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	db.Init()

	server.Init()

}

