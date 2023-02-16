package main

import (
	"bankAPI/internal/server"
	"bankAPI/internal/storage"
)

// Starting point
func main() {
	bankDB := storage.ConnectDB()
	mainStorage := storage.NewStorage(bankDB)
	mainServer := server.NewServer(mainStorage)
	mainServer.Start()
}
