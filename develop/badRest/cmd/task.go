package main

// здесь я буду писать сервер.

import (
	"net/http"

	"../pkg/useCases"
)

func main() {

}

func RunWebServer(){
	mux := http.NewServeMux()
	mux.HandleFunc("/events_for_day", EventsForDay)
	http.ListenAndServe("127.0.0.1", mux)
}

func EventsForDay(w http.ResponseWriter, r *http.Request) {

}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	allEvents := useCases.GetAllEvents()
}
