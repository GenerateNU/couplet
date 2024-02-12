package main

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target ../../internal/api --clean ../../../openapi.yaml

import (
	"context"
	"couplet/internal/controller"
	"couplet/internal/database"
	"couplet/internal/handler"
	"fmt"
	"log/slog"
	"os"

	"log"
	"net/http"

	"couplet/internal/api"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
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

	Port     uint16 `env:"PORT, default=8080"`      // the port for the server to listen on
	LogLevel string `env:"LOG_LEVEL, default=INFO"` // the level of event to log
}

func main() {
	// Load environment variables
	ctx := context.Background()
	var config EnvConfig
	var err error
	if err = envconfig.Process(ctx, &config); err != nil {
		log.Fatalln(err)
	}

	// Display splash screen. Purely cosmetic :)
	logo, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("couplet", pterm.FgMagenta.ToStyle())).Srender()
	pterm.DefaultCenter.Println(logo)
	credit := pterm.DefaultBox.Sprint("Prototype created by " + pterm.Cyan("Generate"))
	pterm.DefaultCenter.Println(credit)

	// Configure slog logger
	logLevel := AsLogLevel(config.LogLevel)
	logger := slog.New(pterm.NewSlogHandler(pterm.DefaultLogger.WithLevel(logLevel)))

	// Connect to the database
	var db *gorm.DB
	if db, err = database.NewDB(config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName, logger); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	if err = database.EnableConnPooling(db); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("server successfully connected to database")

	// Instantiate a controller for business logic
	var c controller.Controller
	if c, err = controller.NewController(db, logger); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Instantiate a handler for serving API requests
	h := handler.NewHandler(c, logger)

	// Instantiate generated server
	var s *api.Server
	if s, err = api.NewServer(h); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("server successfully instantiated and listening", "port", config.Port)

	// Run server indefinitely until an error occurs
	if err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), s); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

// Converts a string to its corresponding log level
func AsLogLevel(logLevel string) pterm.LogLevel {
	switch logLevel {
	case "DEBUG":
		return pterm.LogLevelDebug
	case "INFO":
		return pterm.LogLevelInfo
	case "WARN":
		return pterm.LogLevelWarn
	case "ERROR":
		return pterm.LogLevelError
	default:
		return pterm.LogLevelDisabled
	}
}
