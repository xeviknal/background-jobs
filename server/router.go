package server

import (
	"net/http"

	"github.com/xeviknal/background-jobs/handlers"
)

type Route struct {
	Name       string
	Path       string
	Method     string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func GetRoutes() Routes {
	return Routes{
		{
			Path:       "/",
			Method:     "GET",
			HandleFunc: handlers.Root,
		},
		{
			Path:       "/jobs/create",
			Method:     "POST",
			HandleFunc: handlers.JobsCreateHandler,
		},
	}
}
