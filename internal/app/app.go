package app

import (
	"log"
	"time"
	"net/http"
	"github.com/paulsonlegacy/go-social/internal/models"
	"github.com/go-chi/chi/v5"
)

// App Engline
type Application struct {
	Config Config
	Models models.Models
}

type Config struct {
	ServerAddress string
	DB struct {
		DBURL string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime string
	}
}

// Run method uses engine parameters to start a server
func (app *Application) Run(router *chi.Mux) error {
	
	// Declaring server parameters
	server := &http.Server{
		Addr: app.Config.ServerAddress,
		Handler: router,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Server has started at %s", app.Config.ServerAddress)

	return server.ListenAndServe()
}