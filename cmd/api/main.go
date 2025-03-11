package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/paulsonlegacy/go-env-loader"
)

// VARIABLES

var BASE_PATH string


// FUNCTIONS

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	BASE_PATH = dir // Sets the root directory dynamically
}


// APPLICATION ENTRY POINT

func main() {
	envPath := filepath.Join(BASE_PATH, "internal/env/.env") // ENV file path

	server_address, err := envloader.GetEnv(envPath, "SERVER_ADDRESS")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	configuration := config{
		address: server_address,
	}

	app := &application{
		config: configuration,
	}

	mux := app.mount()

	log.Printf("Server has started at %s", app.config.address)

	log.Fatal(app.run(mux))
}