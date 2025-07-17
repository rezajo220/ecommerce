package main

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	godotenv.Load()

	port := getEnv("SERVER_PORT", "8080")
	readTimeoutSec, _ := strconv.Atoi(getEnv("SERVER_READ_TIMEOUT", "30"))
	writeTimeoutSec, _ := strconv.Atoi(getEnv("SERVER_WRITE_TIMEOUT", "30"))

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USERNAME", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "ecommerce")
	dbSSLMode := getEnv("DB_SSL_MODE", "disable")

	config := &Config{
		Server: ServerConfig{
			Port:         port,
			ReadTimeout:  time.Duration(readTimeoutSec) * time.Second,
			WriteTimeout: time.Duration(writeTimeoutSec) * time.Second,
		},
		Database: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			DBName:   dbName,
			SSLMode:  dbSSLMode,
		},
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
