package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paulsonlegacy/go-social/internal/handlers"
)

// Set up the router and connects handlers to routes.
func SetUpRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// Routes
	r.Route("/api/v1", func(r chi.Router) {
		// Using handler from handlers package
		r.Get("/", handlers.HomeHandler)
	})

	return r
}