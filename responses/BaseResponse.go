package responses

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (br *BaseResponse) SendJson(w http.ResponseWriter, code int) {
	response, _ := json.Marshal(br)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
