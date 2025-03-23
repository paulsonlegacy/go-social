package main

import (
	"log"
	"os"
	"time"
	"path/filepath"
	"net/http"
	envloader "github.com/paulsonlegacy/go-env-loader"
	"github.com/paulsonlegacy/go-social/internal/db"
	"github.com/paulsonlegacy/go-social/internal/models"
	"github.com/paulsonlegacy/go-social/internal/routes"
	"github.com/go-chi/chi/v5"
)


// CONSTANTS

var (
	BASE_PATH string
	ENV_PATH string
	SERVER_ADDRESS string
) 


// STRUCTURES

type Config struct {
	ServerAddress string
	DB struct {
		DBURL string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime string
	}
}

type Application struct {
	Config Config
	Models models.Models
}

func (app *Application) run(router *chi.Mux) error {
	
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


// FUNCTIONS

// init() runs before main function is executed
func init() {
	// init() Usage – Prepares paths and server variables before main()

	// Setting BASE path from this file
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	BASE_PATH = dir // Sets the root directory dynamically


	// Setting env path
	envPath := filepath.Join(BASE_PATH, "internal/config/.env") // ENV file 
	ENV_PATH = envPath
	
	
	// Setting server address
	server_address, err := envloader.GetEnv(ENV_PATH, "SERVER_ADDRESS")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	SERVER_ADDRESS = server_address
}



// Application entry point
func main() {
	// Variables
	db_url, _ := envloader.GetEnv(ENV_PATH, "DBURL")
	
	// Types
	configurations := Config{
		ServerAddress: SERVER_ADDRESS,
		DB: struct {
			DBURL        string
			MaxOpenConns int
			MaxIdleConns int
			MaxIdleTime  string
		}{
			DBURL:        db_url,
			MaxOpenConns: 30,
			MaxIdleConns: 30,
			MaxIdleTime:  "15m",
		},
	}

	// Initializing new DB connection
	DB_CONNECTION, err := db.NewDBConnection(
		configurations.DB.DBURL,
		configurations.DB.MaxOpenConns,
		configurations.DB.MaxIdleConns,
		configurations.DB.MaxIdleTime,
	)

	if err != nil { // Error while initializing DB connection
		log.Panic(err)
	}

	// Initializing App engine
	app := &Application{
		Config: configurations,
		Models:  models.NewModels(DB_CONNECTION),
	}

	// Initializing mux/router
	appRouter := router.SetUpRouter()

	// Running server
	log.Fatal(app.run(appRouter))
}
