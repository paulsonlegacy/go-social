package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paulsonlegacy/go-social/internal/app"
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
		r.Get("/", app.HomeHandler)

		r.Route("/posts", func (r chi.Router) {
			r.Get("/", app.FetchPostsHandler)
			r.Get("/{id}", app.FetchPostHandler)
			r.Post("/create", app.CreatePostHandler)
			r.Patch("/{id}/update", app.UpdatePostHandler)
			r.Delete("/{id}/delete", app.DeletePostHandler)
		})


	})

	return r
}