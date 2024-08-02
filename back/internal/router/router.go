package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"quiz-aventure/internal/config"
	"quiz-aventure/internal/handlers"
)

// NewRouter returns a new router
func NewRouter(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORS.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(c.Handler)

	r.Get("/questions", handlers.GetQuestions)
	r.Post("/check-answer", handlers.CheckAnswer)
	r.Post("/reset-score", handlers.ResetScore)

	return r
}
