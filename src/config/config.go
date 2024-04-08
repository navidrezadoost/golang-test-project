package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Cors     CorsConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}
type ServerConfig struct {
	Port    int
	RunMode string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
}

type CorsConfig struct {
	AllowOrigins string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SslMode  bool
}

type RedisConfig struct {
	Host               string
	Port               int
	Password           string
	Db                 int
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	fmt.Println(cfgPath)
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("error in load config")
	}
	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("error in parse config")
	}
	return cfg
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "./config/config-docker.yml"
	} else if env == "production" {
		return "./config/config-production"
	} else {
		return "../config/config-development.yml"
	}
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(filename string, filetype string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(filetype)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to parse config file, %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}
