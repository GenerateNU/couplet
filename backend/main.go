package main

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --clean openapi.yaml

import (
	"couplet/config"
	"couplet/database"
	"couplet/handler"

	"log"
	"net/http"

	"couplet/api"
)

func main() {
	// Load config
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Connect to database
	db, err := database.ConfigureDB(config)
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.ConnPooling(db); err != nil {
		log.Fatalln(err)
	}

	// Instantiate API request handler and create generated server
	server, err := api.NewServer(handler.NewHandler(db))
	if err != nil {
		log.Fatalln(err)
	}

	// Run server indefinitely until an error occurs
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalln(err)
	}
}
