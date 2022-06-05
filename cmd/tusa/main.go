package main

import "tusa/internal/router"

func main() {

	server := router.SetupServer()
	server.Run(":32958")
}
