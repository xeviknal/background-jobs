package handlers

import (
	"encoding/json"
	"net/http"
)

type responseError struct {
	Error  string `json:"error,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		response, _ = json.Marshal(responseError{Error: err.Error()})
		code = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func Root(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w, http.StatusOK, "Ready to create jobs!")
}
