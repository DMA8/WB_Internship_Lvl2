package handlers

import (
	"encoding/json"
	"log"
	"../models"
)

type errorResponse struct {
	ErrorMesage	string `json:"error"`
}

type resultResponse struct {
	Result []models.Event `json:"result"`
}

type postSuccessResultResponse struct {
	Result string `json:"result"`
}

func makeJSONErrorResponse(msg string) []byte {
	result, err := json.Marshal(errorResponse{msg})
	if err != nil {
		log.Fatal("making error json didn't succeed")
	}
	return result
}

func makeJSONResultResponse(events []models.Event) []byte{
	result, err := json.Marshal(resultResponse{events})
	if err != nil {
		log.Fatal("making resultResponse json didn't succeed")
	}
	return result
}

func makeJSONPostSuccessResultResponse(msg string) []byte{
	result, err := json.Marshal(postSuccessResultResponse{msg})
	if err != nil {
		log.Fatal("making postResponse json didn't succeed")
	}
	return result
}