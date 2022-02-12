package main

import (
	"net/http"
	
	"../internal/handlers"
	"../internal/storage"
)
//TODO: защитить storage мьютексами
func main() {
	storage := storage.NewEventStorage()
	handler := handlers.NewHandler(storage)

	mux := handler.InitRoutes()
	http.ListenAndServe(":8080", mux)
}