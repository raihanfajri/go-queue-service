package pubctrl

import (
	"encoding/json"
	"net/http"

	"../helpers"
	"../responses"
	pubservice "../services"
)

type PubController struct{}

func (p *PubController) Handle(w http.ResponseWriter, r *http.Request) {
	response := &responses.BaseResponse{}
	code := 200

	var param interface{}

	json.NewDecoder(r.Body).Decode(&param)

	message, _ := json.Marshal(param)
	url := r.URL.Path

	pubService := &pubservice.PubService{}

	sentMessage, err := pubService.PublishMessage(string(message), url)

	response.Message = sentMessage

	if err != nil {
		response.Success = false
		code = 500
		response.Message = "Internal Server Error"
		helpers.LogError(sentMessage, err)
	}

	response.SendJson(w, code)
}
