package main

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --clean ../openapi.yaml

import (
	"couplet/database"
	"couplet/handler"
	"fmt"
	"os"

	"log"
	"net/http"

	"couplet/api"
)

func main() {
	// Load environment variables
	dbHost, envSet := os.LookupEnv("DB_HOST")
	if !envSet {
		log.Fatalln("database host not specified")
	}

	dbPort, envSet := os.LookupEnv("DB_PORT")
	if !envSet {
		log.Fatalln("database port not specified")
	}

	dbUser, envSet := os.LookupEnv("DB_USER")
	if !envSet {
		log.Fatalln("database username not specified")
	}

	dbPassword, envSet := os.LookupEnv("DB_PASSWORD")
	if !envSet {
		log.Fatalln("database password not specified")
	}

	dbName, envSet := os.LookupEnv("DB_NAME")
	if !envSet {
		log.Fatalln("database name not specified")
	}

	port, envSet := os.LookupEnv("PORT")
	if !envSet {
		log.Fatalln("port not specified")
	}

	// Connect to database
	db, err := database.ConfigureDB(dbHost, dbPort, dbUser, dbPassword, dbName)
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
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), server); err != nil {
		log.Fatalln(err)
	}
}
