// Package config provides functionality for loading and parsing application configurations
// from various sources such as environment variables and configuration files.
package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

// Config represents the application configuration structure.
type Config struct {
	SERVER  ServerConfig
	LOGGER  LoggerConfig
	CORS    CorsConfig
	MONGODB MongoDbConfig
	REDIS   RedisConfig
}

// ServerConfig represents the configuration for the server.
type ServerConfig struct {
	PORT              int
	CONTROLLER_PREFIX string
	MODE              string
}

// LoggerConfig represents the configuration for logging.
type LoggerConfig struct {
	FILE_PATH   string
	ENCODING    string
	LEVEL       string
	MAX_SIZE    int
	MAX_AGE     int
	MAX_BACKUPS int
	COMPRESS    bool
	LOGGER      string
}

// CorsConfig represents the configuration for Cross-Origin Resource Sharing (CORS).
type CorsConfig struct {
	ALLOW_ORIGINS string
}

// MongoDbConfig represents the configuration for MongoDB.
type MongoDbConfig struct {
	HOST                 string
	PORT                 int
	USER                 string
	PASSWORD             string
	DATABASE_NAME        string
	SSL_MODE             string
	TIMEZONE             string
	MAX_IDLE_CONNECT     int
	MAX_OPEN_CONNECT     int
	CONNECT_MAX_LIFETIME time.Duration
}

// RedisConfig represents the configuration for Redis.
type RedisConfig struct {
	HOST                 string
	POST                 int
	PASSWORD             string
	DB                   int
	DIAL_TIMEOUT         time.Duration
	READ_TIMEOUT         time.Duration
	WRITE_TIMEOUT        time.Duration
	IDLE_CHECK_FREQUENCY time.Duration
	IDLE_TIMEOUT         time.Duration
	POOL_SIZE            int
	POOL_TIMEOUT         time.Duration
}

// GetConfig loads and returns the application configuration.
func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("error in loading config: %v", err)
	}
	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("error in parsing config: %v", err)
	}
	return cfg
}

// getConfigPath returns the path to the configuration file based on the environment.
func getConfigPath(env string) string {
	if env == "docker" {
		return "./config/config-docker.yml"
	} else if env == "production" {
		return "./config/config-production.yml"
	} else {
		return "./handlers/config/config-development.yml"
	}
}

// ParseConfig parses the configuration from the given Viper instance.
func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// LoadConfig loads the configuration from the specified file and type.
func LoadConfig(filename string, filetype string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(filetype)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		log.Printf("unable to parse config file: %v", err)
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}
