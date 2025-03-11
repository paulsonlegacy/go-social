package main

import (
	"log"
)

func main() {
	configuration := config{
		address: ":8080",
	}
	app := &application{
		config: configuration,
	}

	mux := app.mount()

	log.Printf("Server has started at %s", app.config.address)

	log.Fatal(app.run(mux))
}