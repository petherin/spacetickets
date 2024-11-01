package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	dBUsernameEnvVar          = "DB_USERNAME"
	dbPasswordEnvVar          = "DB_PASSWORD"
	dbNameEnvVar              = "DB_NAME"
	dbHostEnvVar              = "DB_HOST"
	maxOpenConnsEnvVar        = "MAX_OPEN_CONNS"
	maxIdleConnsEnvVar        = "MAX_IDLE_CONNS"
	connMaxLifetimeSecsEnvVar = "CONN_MAX_LIFETIME_SECS"
	connRetriesEnvVar         = "DB_CONN_RETRIES"
	retryIntervalEnvVar       = "DB_CONN_RETRY_INTERVAL_SECS"
	apiPortEnvVar             = "API_PORT"
	swaggerPortEnvVar         = "SWAGGER_PORT"
)

type Config struct {
	DBUsername              string
	DBPassword              string
	DBName                  string
	DBHost                  string
	MaxOpenConns            int
	MaxIdleConns            int
	ConnMaxLifetimeSecs     int
	DBConnRetries           int
	DBConnRetryIntervalSecs int
	APIPort                 string
	SwaggerPort             string
}

// Get retrieves config from environment variables.
func Get() (Config, error) {
	username := os.Getenv(dBUsernameEnvVar)
	if len(username) == 0 {
		return Config{}, fmt.Errorf("unrecognised value for environment variable %s", dBUsernameEnvVar)
	}

	password := os.Getenv(dbPasswordEnvVar)
	if len(password) == 0 {
		return Config{}, fmt.Errorf("unrecognised value for environment variable %s", dbPasswordEnvVar)
	}

	name := os.Getenv(dbNameEnvVar)
	if len(name) == 0 {
		return Config{}, fmt.Errorf("unrecognised value for environment variable %s", dbNameEnvVar)
	}

	host := os.Getenv(dbHostEnvVar)
	if len(host) == 0 {
		return Config{}, fmt.Errorf("unrecognised value for environment variable %s", dbHostEnvVar)
	}

	openConns, err := GetEnvVarInt(maxOpenConnsEnvVar)
	if err != nil {
		return Config{}, err
	}

	idleConns, err := GetEnvVarInt(maxIdleConnsEnvVar)
	if err != nil {
		return Config{}, err
	}

	connLifetime, err := GetEnvVarInt(connMaxLifetimeSecsEnvVar)
	if err != nil {
		return Config{}, err
	}

	retries, err := GetEnvVarInt(connRetriesEnvVar)
	if err != nil {
		return Config{}, err
	}

	interval, err := GetEnvVarInt(retryIntervalEnvVar)
	if err != nil {
		return Config{}, err
	}

	port := os.Getenv(apiPortEnvVar)
	if len(apiPortEnvVar) == 0 {
		return Config{}, fmt.Errorf("unrecognised value for environment variable %s", apiPortEnvVar)
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	swagPort := os.Getenv(swaggerPortEnvVar)
	if len(swagPort) == 0 {
		return Config{}, fmt.Errorf("unrecognised value for environment variable %s", swaggerPortEnvVar)
	}
	if !strings.HasPrefix(swagPort, ":") {
		swagPort = ":" + swagPort
	}

	cfg := Config{
		DBUsername:              username,
		DBPassword:              password,
		DBName:                  name,
		DBHost:                  host,
		MaxOpenConns:            openConns,
		MaxIdleConns:            idleConns,
		ConnMaxLifetimeSecs:     connLifetime,
		DBConnRetries:           retries,
		DBConnRetryIntervalSecs: interval,
		APIPort:                 port,
		SwaggerPort:             swagPort,
	}

	log.Println("Config loaded from environment variables")

	return cfg, nil
}

func GetEnvVarInt(name string) (int, error) {
	valueStr := os.Getenv(name)
	if len(valueStr) == 0 {
		return 0, fmt.Errorf("unrecognised value for environment variable %s", name)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("can't parse environment variable %s", name)
	}

	return value, nil
}
