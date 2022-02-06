package handlers

import (
	"net/http"
	"net/url"
	"time"

	"../models"
)

func getDate(u *url.URL) (time.Time, error){
	date := u.Query().Get("date")
	return time.Parse(models.TimeFormat, date)
}

func (h *Handler) GetEventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	date, err := getDate(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongDateFormat))
		return
	}
	events := h.storage.GetEventsForDay(date)
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONResultResponse(events))
}

func (h *Handler) GetEventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	date, err := getDate(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongDateFormat))
		return
	}
	events := h.storage.GetEventsForWeek(date)
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONResultResponse(events))
}

func (h *Handler) GetEventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	date, err := getDate(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongDateFormat))
		return
	}
	events := h.storage.GetEventsForMonth(date)
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONResultResponse(events))
}