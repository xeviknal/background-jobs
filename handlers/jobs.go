package handlers

import "net/http"

func JobsCreateHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w, http.StatusOK, "Hello world!")
}
