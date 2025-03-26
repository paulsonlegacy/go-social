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
		r.Post("/posts/create", func(w http.ResponseWriter, r *http.Request) {
			handlers.CreatePostHandler(w, r, app)
		})
		r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
			handlers.FetchPostsHandler(w, r, app)
		})
		r.Get("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
			handlers.FetchPostHandler(w, r, app)
		})
		r.Patch("/posts/{id}/update", func(w http.ResponseWriter, r *http.Request) {
			handlers.UpdatePostHandler(w, r, app)
		})
		r.Delete("/posts/{id}/delete", func(w http.ResponseWriter, r *http.Request) {
			handlers.DeletePostHandler(w, r, app)
		})
	})

	return r
}