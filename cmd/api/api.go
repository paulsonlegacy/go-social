package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paulsonlegacy/go-social/internal/model"
)

type application struct {
	config config
	model model.Models
}

type config struct {
	server_address string
	db dbConfig
}

type dbConfig struct {
	dbhost string
	dbport int
	dbuser string
	dbpass string
	dburl string
	maxOpenConnections int
	maxIdleConnections int
	maxIdleTime string
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
		Addr: app.config.server_address,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	return server.ListenAndServe()
}