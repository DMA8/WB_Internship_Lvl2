package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

func deserializateEvent(r *http.Request) (models.Event, error){
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event) // оптимальнее, чем Unmarshal
	defer r.Body.Close()
	return event, err
}

func (h *Handler) createEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := deserializateEvent(r)//
	fmt.Println(event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(makeJSONErrorResponse(msgWrongInputData))
		return
	}

	_, err =  h.storage.CreateEvent(event)
	if err != nil{
		w.WriteHeader(503)
		w.Write(makeJSONErrorResponse(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(makeJSONPostSuccessResultResponse(msgCreateSuccess))
}

func (h *Handler) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := deserializateEvent(r)
	fmt.Println(h.storage.Events)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(makeJSONErrorResponse(msgWrongInputData))
		return
	}

	err =  h.storage.DeleteEvent(event.EventID)
	if err != nil{
		w.WriteHeader(503)
		w.Write(makeJSONErrorResponse(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(makeJSONPostSuccessResultResponse(msgDeleteSuccess))
}

func (h *Handler) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := deserializateEvent(r)
	fmt.Println(event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(makeJSONErrorResponse(msgWrongInputData))
		return
	}

	_, err =  h.storage.UpdateEvent(event.EventID, event)
	if err != nil{
		w.WriteHeader(503)
		w.Write(makeJSONErrorResponse(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(makeJSONPostSuccessResultResponse(msgUpdatedSuccess))
}