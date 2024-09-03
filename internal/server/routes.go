package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)

	// User routes
	r.Route("/users", func(r chi.Router) {
		r.Post("/", s.userHandler.Create)
		r.Get("/", s.userHandler.List)
		r.Get("/{id}", s.userHandler.Get)
		r.Put("/{id}", s.userHandler.Update)
		r.Delete("/{id}", s.userHandler.Delete)
	})

	// Channel routes
	r.Route("/channels", func(r chi.Router) {
		r.Post("/", s.channelHandler.Create)
		r.Get("/", s.channelHandler.List)
		r.Get("/{id}", s.channelHandler.Get)
		r.Put("/{id}", s.channelHandler.Update)
		r.Delete("/{id}", s.channelHandler.Delete)
		r.Post("/{id}/users", s.channelHandler.AddUser)
		r.Delete("/{id}/users/{userId}", s.channelHandler.RemoveUser)
		r.Get("/{id}/users", s.channelHandler.ListUsers)
	})

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
