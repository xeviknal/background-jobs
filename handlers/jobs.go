package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/xeviknal/background-jobs/database"
	"github.com/xeviknal/background-jobs/models"
)

func JobsCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	objectIdRaw := vars["object_id"]
	objectId, err := strconv.ParseInt(objectIdRaw, 10, 64)
	if err != nil {
		RespondWithJson(w, http.StatusUnprocessableEntity, fmt.Sprintf("object_id must be an integer: %v", err))
		return
	}

	db := database.GetDb()
	if db == nil {
		RespondWithJson(w, http.StatusInternalServerError, "database not available")
		return
	}

	// Create the object: sleep between 0 and 10s
	job := models.NewJob(objectId)

	// Save it
	if err := db.Insert(&job); err != nil {
		RespondWithJson(w, http.StatusInternalServerError, fmt.Sprintf("unable to create the new post: %v", err))
		return
	}

	RespondWithJson(w, http.StatusOK, job)
}
