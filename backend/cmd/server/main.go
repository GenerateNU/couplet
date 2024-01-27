package main

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target ../../internal/api --clean ../../../openapi.yaml

import (
	"context"
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/handler"
	"fmt"

	"log"
	"net/http"

	"couplet/internal/api"

	"github.com/sethvargo/go-envconfig"
	"gorm.io/gorm"
)

// Environment variables used to configure the server
type EnvConfig struct {
	DbHost     string `env:"DB_HOST, required"`     // the database host to connect to
	DbPort     uint16 `env:"DB_PORT, required"`     // the database port to connect to
	DbUser     string `env:"DB_USER, required"`     // the user to connect to the database with
	DbPassword string `env:"DB_PASSWORD, required"` // the password to connect to the database with
	DbName     string `env:"DB_NAME, required"`     // the name of the database to connect to

	Port uint16 `env:"PORT, default=8080"` // the port for the server to listen on
}

func main() {
	// Load environment variables
	ctx := context.Background()
	var config EnvConfig
	var err error
	if err = envconfig.Process(ctx, &config); err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	var db *gorm.DB
	if db, err = database.NewDb(config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName); err != nil {
		log.Fatalln(err)
	}
	if err = database.EnableConnPooling(db); err != nil {
		log.Fatalln(err)
	}

	// Instantiate a controller for business logic
	var c controller.Controller
	if c, err = controller.NewController(db); err != nil {
		log.Fatalln(err)
	}

	// Instantiate a handler for serving API requests
	h := handler.NewHandler(c)

	// Instantiate generated server
	var s *api.Server
	if s, err = api.NewServer(h); err != nil {
		log.Fatalln(err)
	}

	// Run server indefinitely until an error occurs
	if err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), s); err != nil {
		log.Fatalln(err)
	}
}
