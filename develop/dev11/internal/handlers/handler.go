package handlers

import (
	"log"
	"net/http"

	"../storage"
)

const(
	MsgWrongDateFormat = "wrong date format"
	MsgMethodNotAllowed = "method not allowed"
	MsgWrongInputData = "wrong input data"
	MsgCreateSuccess = "msg created"
	MsgDeleteSuccess = "event deleted"
	MsgUpdatedSuccess = "event updated"
)

type Handler struct {
	storage *storage.EventStorage
}

func NewHandler(stor *storage.EventStorage) *Handler{
	return &Handler{
		storage:  stor,
	}
}

func (h *Handler)InitRoutes() *http.ServeMux {
	mux := http.ServeMux{}
	mux.HandleFunc("/create_event", Log(CheckHTTPVerb(http.HandlerFunc(h.CreateEventHandler), http.MethodPost)))
	mux.HandleFunc("/delete_event", Log(CheckHTTPVerb(http.HandlerFunc(h.DeleteEventHandler), http.MethodPost)))
	mux.HandleFunc("/update_event", Log(CheckHTTPVerb(http.HandlerFunc(h.UpdateEventHandler), http.MethodPost)))

	mux.HandleFunc("/events_for_day", Log(CheckHTTPVerb(http.HandlerFunc(h.GetEventsForDayHandler), http.MethodGet)))
	mux.HandleFunc("/events_for_week", Log(CheckHTTPVerb(http.HandlerFunc(h.GetEventsForWeekHandler), http.MethodGet)))
	mux.HandleFunc("/events_for_month", Log(CheckHTTPVerb(http.HandlerFunc(h.GetEventsForMonthHandler), http.MethodGet)))

	return &mux
}

func CheckHTTPVerb(handler http.Handler, method string) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method{
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(500)
			w.Write(MakeJSONErrorResponse(MsgMethodNotAllowed))
		}
	}
}

func Log(handler http.Handler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		handler.ServeHTTP(w, r)
		log.Println(r.RequestURI, "is handled")
	}
}
