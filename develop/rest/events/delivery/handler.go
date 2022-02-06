package delivery

import (
	"encoding/json"
	"net/http"

	"../../usecases"
)



func GetAllEvents( w http.ResponseWriter, r *http.Request){
	allEvents, err := h.useCase.GetAllEvents()
	if err != nil{
		// ошибка бизнес логики
	}
	ans, err := json.Marshal(allEvents)
	if err != nil{
		// ошибка бизнес логики
	}
	w.Write(ans)
}