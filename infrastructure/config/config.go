package config

import (
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	RepositoriesConfig
	WebConfig
}

type RepositoriesConfig struct {
	SQLConfig
}

type SQLConfig struct {
	Host     string
	UserName string
	Password string
	DbName   string
	Port     int
	PoolSize int
}

type WebConfig struct {
	ApplicationID string
	Address       string
	AuthKey       string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		RepositoriesConfig: RepositoriesConfig{
			SQLConfig: SQLConfig{
				Host:     getEnv("DB_HOST"),
				UserName: getEnv("DB_USER"),
				Password: getEnv("DB_PASS"),
				DbName:   getEnv("DB_NAME"),
				Port:     getIntEnv("DB_PORT"),
				PoolSize: getIntEnv("DB_MAX_POOL_SIZE"),
			},
		},
		WebConfig: WebConfig{
			ApplicationID: getEnv("APPLICATION_ID"),
			Address:       getEnv("PORT"),
			AuthKey:       getEnv("AUTH_KEY"),
		},
	}
}

func getEnv(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		log.Fatalf("Missing environment variable: '%s' must be set", envName)
	}

	return value
}

func getIntEnv(envName string) int {
	value := getEnv(envName)
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Wrong environment variable type. Expected '%s' of type int", envName)
	}

	return result
}
