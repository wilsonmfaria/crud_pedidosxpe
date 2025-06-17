package view

import (
	"encoding/json"
	"net/http"
)

// RespondJSON envia uma resposta HTTP com conteúdo JSON
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Erro ao gerar JSON de resposta")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// RespondError envia um erro em formato JSON padronizado
func RespondError(w http.ResponseWriter, status int, mensagem string) {
	RespondJSON(w, status, map[string]string{"erro": mensagem})
}
