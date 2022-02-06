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

func (h *Handler) getEventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	date, err := getDate(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(makeJSONErrorResponse(msgWrongDateFormat))
		return
	}
	events := h.storage.GetEventsForDay(date)
	w.WriteHeader(http.StatusOK)
	w.Write(makeJSONResultResponse(events))
}

func (h *Handler) getEventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	date, err := getDate(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(makeJSONErrorResponse(msgWrongDateFormat))
		return
	}
	events := h.storage.GetEventsForWeek(date)
	w.WriteHeader(http.StatusOK)
	w.Write(makeJSONResultResponse(events))
}

func (h *Handler) getEventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	date, err := getDate(r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(makeJSONErrorResponse(msgWrongDateFormat))
		return
	}
	events := h.storage.GetEventsForMonth(date)
	w.WriteHeader(http.StatusOK)
	w.Write(makeJSONResultResponse(events))
}