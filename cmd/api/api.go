package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	address string
}

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/", app.homeHandler)
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	

	server := &http.Server{
		Addr: app.config.address,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	return server.ListenAndServe()
}