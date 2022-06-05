package main

import (
	"os"
	"tusa/internal/router"
)

func main() {

	server := router.SetupServer()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":32958"
	} else {
		port = ":" + port
	}
	server.Run(port)
}
