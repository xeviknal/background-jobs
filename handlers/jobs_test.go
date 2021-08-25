package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestJobsCreateHandler(t *testing.T) {
	asserter := assert.New(t)

	s := setupEndpoint(t)
	defer s.Close()

	resp, err := http.Get(fmt.Sprintf("%s/jobs/create", s.URL))
	if err != nil {
		asserter.Fail(err.Error())
	}

	asserter.Equal(resp.StatusCode, http.StatusOK)
}

func setupEndpoint(t *testing.T) *httptest.Server {
	mr := mux.NewRouter()
	mr.HandleFunc("/jobs/create", JobsCreateHandler)
	return httptest.NewServer(mr)
}
