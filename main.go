package main

import (
	"github.com/ElizeuS/gouser/database"
	"github.com/ElizeuS/gouser/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
