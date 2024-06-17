package main

import "github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server"

// main is the entry point of the Go program.
//
// It calls the StartServer function from the server package to start the server.
// If an error occurs during the server startup, it panics.
func main() {
	if err := server.StartServer(); err != nil {
		panic(err)
	}
}
