package router

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paulsonlegacy/go-social/internal/app"
	"github.com/paulsonlegacy/go-social/internal/handlers"
)

// Set up the router and connects handlers to routes.
func SetUpRouter(app *app.Application) *chi.Mux {
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
		r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
			handlers.CreatePostHandler(w, r, app)
		})
		r.Post("/posts/create", func(w http.ResponseWriter, r *http.Request) {
			handlers.CreatePostHandler(w, r, app)
		})
	})

	return r
}