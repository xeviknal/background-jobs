package server

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/xeviknal/background-jobs/database"
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer() *Server {
	// Prepare the available routes
	r := mux.NewRouter()
	for _, route := range GetRoutes() {
		r.HandleFunc(route.Path, route.HandleFunc).Methods(route.Method).Name(route.Name)
	}

	server := &http.Server{
		Addr:         "0.0.0.0:8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}

	return &Server{
		httpServer: server,
		router:     r,
	}
}

// Start method starts all the necessary services to have the server fully working
func (s *Server) Start() {
	// Starting a seed for randoms
	rand.Seed(time.Now().UnixNano())

	// Init Database
	if db := database.GetDb(); db == nil {
		log.Fatalln("Errors initializing database systems")
	}

	// Starting the server in background
	go func() {
		log.Printf("Starting server at %s", s.httpServer.Addr)
		err := s.httpServer.ListenAndServe()
		if err != nil {
			log.Printf("Error while staring the server: %v", err)
		}
	}()
}

// Stop method closes all the necessary to gracefully shutdown the server
func (s *Server) Stop() {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Shutting down the server
	s.httpServer.Shutdown(ctx)
}
