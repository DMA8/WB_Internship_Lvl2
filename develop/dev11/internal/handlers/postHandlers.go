package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

func DeserializateEvent(r *http.Request) (models.Event, error){
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event) // оптимальнее, чем Unmarshal
	return event, err
}

func (h *Handler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := DeserializateEvent(r)//
	fmt.Println(event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongInputData))
		return
	}

	_, err =  h.storage.CreateEvent(event)
	if err != nil{
		w.WriteHeader(503)
		w.Write(MakeJSONErrorResponse(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONPostSuccessResultResponse(MsgCreateSuccess))
}

func (h *Handler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := DeserializateEvent(r)
	fmt.Println(h.storage.Events)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongInputData))
		return
	}

	err =  h.storage.DeleteEvent(event.EventID)
	if err != nil{
		w.WriteHeader(503)
		w.Write(MakeJSONErrorResponse(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONPostSuccessResultResponse(MsgDeleteSuccess))
}

func (h *Handler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := DeserializateEvent(r)
	fmt.Println(event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(MakeJSONErrorResponse(MsgWrongInputData))
		return
	}

	_, err =  h.storage.UpdateEvent(event.EventID, event)
	if err != nil{
		w.WriteHeader(503)
		w.Write(MakeJSONErrorResponse(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(MakeJSONPostSuccessResultResponse(MsgUpdatedSuccess))
}