package httphandler

import (
	"encoding/json"
	"net/http"
)

func RespondError(mensagem string, resposta map[string]any, w http.ResponseWriter) {
	resposta["Sucess"] = false
	resposta["Message"] = mensagem
	respostaByte, _ := json.Marshal(resposta)
	w.Header().Add("Content-Type", "application/json")
	w.Write(respostaByte)
}

func RespondSucess(resposta map[string]any, w http.ResponseWriter) {
	respostaByte, err := json.Marshal(resposta)
	if err != nil {
		RespondError(err.Error(), resposta, w)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(respostaByte)
}
