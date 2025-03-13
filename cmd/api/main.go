package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/paulsonlegacy/go-env-loader"
	"github.com/paulsonlegacy/go-social/internal/database"
	"github.com/paulsonlegacy/go-social/internal/model"
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

	DBURL, _ := envloader.GetEnv(envPath, "DBURL")

	configurations := config{
		server_address: server_address,
		db: dbConfig{
			dburl: DBURL,
			maxOpenConnections: 30,
			maxIdleConnections: 30,
			maxIdleTime: "15mins",
		},
	}

	db, err := db.New(
		configurations.db.dburl,
		configurations.db.maxOpenConnections,
		configurations.db.maxIdleConnections,
		configurations.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	model := model.NewModel(db)

	app := &application{
		config: configurations,
		model: model,
	}

	mux := app.mount()

	log.Printf("Server has started at %s", app.config.server_address)

	log.Fatal(app.run(mux))
}