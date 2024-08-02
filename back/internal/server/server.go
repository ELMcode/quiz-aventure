package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"quiz-aventure/internal/config"
	"quiz-aventure/internal/router"
)

type Server struct {
	cfg    *config.Config
	router *chi.Mux
}

// NewServer returns a new server
func NewServer(cfg *config.Config) *Server {
	r := router.NewRouter(cfg)
	return &Server{
		cfg:    cfg,
		router: r,
	}
}

// Run starts the server
func (s *Server) Run() error {
	log.Printf("Server is running at %s", s.cfg.Server.Address)
	return http.ListenAndServe(s.cfg.Server.Address, s.router)
}
