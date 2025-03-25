package main

import (
	"log"
	"os"
	"path/filepath"
	
	envloader "github.com/paulsonlegacy/go-env-loader"
	"github.com/paulsonlegacy/go-social/internal/app"
	"github.com/paulsonlegacy/go-social/internal/db"
	"github.com/paulsonlegacy/go-social/internal/models"
	"github.com/paulsonlegacy/go-social/internal/routes"
)


// CONSTANTS

var (
	BASE_PATH string
	ENV_PATH string
	SERVER_ADDRESS string
) 


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
	configurations := app.Config{
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
	var app *app.Application = &app.Application{
		Config: configurations,
		Models:  models.NewModels(DB_CONNECTION),
	}

	// Initializing mux/router
	appRouter := router.SetUpRouter(app)

	// Running server
	log.Fatal(app.Run(appRouter))
}
