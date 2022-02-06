package main

import (
	"net/http"
	"../events/delivery"
)

func main() {
	
}

func WebServ(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", delivery.Handler.GetAllEvents)
	http.ListenAndServe("127.0.0.1:2000", mux)
}