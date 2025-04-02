package main

import (
	"fmt"
	"os"

	"github.com/OleksandrZhurba-san/ichgram-server/server"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := server.Init()

	fmt.Printf("listening on port: %s", port)

	r.Run(":" + port)

}