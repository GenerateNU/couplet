// Read configurations from a YAML file
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Backend  BackendConfig  `yaml:"backend"`
	Database DatabaseConfig `yaml:"database"`
}

type BackendConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         uint16 `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"databaseName"`
	RequireSSL   bool   `yaml:"requireSsl"`
}

func (s *DatabaseConfig) String() string {
	sslMode := "disable"
	if s.RequireSSL {
		sslMode = "require"
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s dbname=%s",
		s.Host, s.Port, s.Username, s.Password, sslMode, s.DatabaseName)
}

const (
	EnvironmentDev string = "dev"
)

// Loads the YAML config file specified by the environment variable ENV.
// Defaults to "dev"
func GetConfig() (Config, error) {
	var config Config

	// Find config directory
	ex, err := os.Executable()
	if err != nil {
		return config, err
	}
	viper := viper.New()
	viper.AddConfigPath(fmt.Sprintf("%s/../config", filepath.Dir(ex)))
	viper.SetConfigType("yaml")

	// Select the specified config file
	env := EnvironmentDev
	if envValue, envSet := os.LookupEnv("ENV"); envSet {
		env = envValue
	}
	viper.SetConfigName(env)

	// Load the config file
	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("failed to read %s configuration: %w", env, err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}
	return config, nil
}
