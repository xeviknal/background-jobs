package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/xeviknal/background-commons/database"
	"github.com/xeviknal/background-commons/models"
)

func TestJobsCreateHandler(t *testing.T) {
	asserter := assert.New(t)

	s := setupEndpoint("/objects/{object_id}/jobs/create", JobsCreateHandler)
	defer s.Close()

	resp, err := http.Get(fmt.Sprintf("%s/objects/%d/jobs/create", s.URL, 10))
	if err != nil {
		asserter.Fail(err.Error())
	}

	asserter.Equal(resp.StatusCode, http.StatusOK)
}

func setupEndpoint(route string, handler http.HandlerFunc) *httptest.Server {
	mr := mux.NewRouter()
	mr.HandleFunc(route, handler)
	return httptest.NewServer(mr)
}

func TestJobsGetSuccess(t *testing.T) {
	asserter := assert.New(t)

	// Creating the Job we want to retrieve
	db := database.GetDb()
	job := models.NewJob(10)
	db.Insert(&job)

	s := setupEndpoint("/jobs/{id}", JobsGetHandler)
	defer s.Close()

	resp, err := http.Get(fmt.Sprintf("%s/jobs/%d", s.URL, job.Id))
	if err != nil {
		asserter.Fail(err.Error())
	}

	asserter.Equal(resp.StatusCode, http.StatusOK)
}

func TestJobsGetNotFound(t *testing.T) {
	asserter := assert.New(t)

	s := setupEndpoint("/jobs/{id}", JobsGetHandler)
	defer s.Close()

	resp, err := http.Get(fmt.Sprintf("%s/jobs/%d", s.URL, 10))
	if err != nil {
		asserter.Fail(err.Error())
	}

	asserter.Equal(resp.StatusCode, http.StatusNotFound)
}
