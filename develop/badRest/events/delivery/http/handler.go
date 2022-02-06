package handler

import (
	"../../usecases/usecases"
)

type Handler struct{
	useCase	usecases.EventUseCase
}

func NewHandler(useCase usecases.EventUseCase) *Handler{
	return &Handler{
		useCase: useCase,
	}
}