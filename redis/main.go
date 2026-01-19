package main

import (
	"log"
	"net"
	"os"
	store "redis/storage"
)

func main() {
	log.Println("Starting a TCP server at port:", 6379)
	dataStore := store.InitializeStore()

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer listener.Close()
	listenConnections(listener, dataStore)
}
