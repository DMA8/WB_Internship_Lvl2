package handlers

import (
	"log"
	"net/http"

	"../storage"
)

const(
	msgWrongDateFormat = "wrong date format"
	msgMethodNotAllowed = "method not allowed"
	msgWrongInputData = "wrong input data"
	msgCreateSuccess = "msg created"
	msgDeleteSuccess = "event deleted"
	msgUpdatedSuccess = "event updated"
)
//Handler cooperates with storage and outer http request
type Handler struct {
	storage *storage.EventStorage
}

//NewHandler is a constructor for Handler
func NewHandler(stor *storage.EventStorage) *Handler{
	return &Handler{
		storage:  stor,
	}
}

//InitRoutes sets up routes for web serv
func (h *Handler)InitRoutes() *http.ServeMux {
	mux := http.ServeMux{}
	mux.HandleFunc("/create_event", logging(checkHTTPVerb(http.HandlerFunc(h.createEventHandler), http.MethodPost)))
	mux.HandleFunc("/delete_event", logging(checkHTTPVerb(http.HandlerFunc(h.deleteEventHandler), http.MethodPost)))
	mux.HandleFunc("/update_event", logging(checkHTTPVerb(http.HandlerFunc(h.updateEventHandler), http.MethodPost)))

	mux.HandleFunc("/events_for_day", logging(checkHTTPVerb(http.HandlerFunc(h.getEventsForDayHandler), http.MethodGet)))
	mux.HandleFunc("/events_for_week", logging(checkHTTPVerb(http.HandlerFunc(h.getEventsForWeekHandler), http.MethodGet)))
	mux.HandleFunc("/events_for_month", logging(checkHTTPVerb(http.HandlerFunc(h.getEventsForMonthHandler), http.MethodGet)))

	return &mux
}

func checkHTTPVerb(handler http.Handler, method string) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method{
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(500)
			w.Write(makeJSONErrorResponse(msgMethodNotAllowed))
		}
	}
}

func logging(handler http.Handler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		handler.ServeHTTP(w, r)
		log.Println(r.RequestURI, "is handled")
	}
}
