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
