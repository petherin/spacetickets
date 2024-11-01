package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	dBUsernameEnvVar              = "DB_USERNAME"
	dbPasswordEnvVar              = "DB_PASSWORD"
	dbNameEnvVar                  = "DB_NAME"
	dbHostEnvVar                  = "DB_HOST"
	dbMaxOpenConnsEnvVar          = "DB_MAX_OPEN_CONNS"
	dbMaxIdleConnsEnvVar          = "DB_MAX_IDLE_CONNS"
	dbConnMaxLifetimeSecsEnvVar   = "DB_CONN_MAX_LIFETIME_SECS"
	dbConnRetriesEnvVar           = "DB_CONN_RETRIES"
	dbRetryIntervalEnvVar         = "DB_CONN_RETRY_INTERVAL_SECS"
	apiPortEnvVar                 = "API_PORT"
	swaggerPortEnvVar             = "SWAGGER_PORT"
	httpTimeoutEnvVar             = "HTTP_TIMEOUT_SECS"
	maxIdleConnsEnvVar            = "MAX_IDLE_CONNS"
	maxConnsPerHostEnvVar         = "MAX_CONNS_PER_HOST"
	idleConnTimeoutSecsEnvVar     = "IDLE_CONN_TIMEOUT_SECS"
	dialerTimeoutSecsEnvVar       = "DIALER_TIMEOUT_SECS"
	dialerKeepAliveSecsEnvVar     = "DIALER_KEEP_ALIVE_SECS"
	tlsHandshakeTimeoutSecsEnvVar = "TLS_HANDSHAKE_TIMEOUT_SECS"
	disableKeepAlivesEnvVar       = "DISABLE_KEEP_ALIVES"
	spaceXAPIEndpointEnvVar       = "SPACEX_API_ENDPOINT"
)

type Config struct {
	DBUsername              string
	DBPassword              string
	DBName                  string
	DBHost                  string
	DBMaxOpenConns          int
	DBMaxIdleConns          int
	DBConnMaxLifetimeSecs   int
	DBConnRetries           int
	DBConnRetryIntervalSecs int
	APIPort                 string
	SwaggerPort             string
	HTTPTimeout             int
	MaxIdleConns            int
	MaxConnsPerHost         int
	IdleConnTimeoutSecs     int
	DialerTimeoutSecs       int
	DialerKeepAliveSecs     int
	TLSHandshakeTimeoutSecs int
	DisableKeepAlives       bool
	SpaceXAPIEndpoint       string
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

	openConns, err := getEnvVarInt(dbMaxOpenConnsEnvVar)
	if err != nil {
		return Config{}, err
	}

	idleConns, err := getEnvVarInt(dbMaxIdleConnsEnvVar)
	if err != nil {
		return Config{}, err
	}

	connLifetime, err := getEnvVarInt(dbConnMaxLifetimeSecsEnvVar)
	if err != nil {
		return Config{}, err
	}

	retries, err := getEnvVarInt(dbConnRetriesEnvVar)
	if err != nil {
		return Config{}, err
	}

	interval, err := getEnvVarInt(dbRetryIntervalEnvVar)
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

	httpTimeout, err := getEnvVarInt(httpTimeoutEnvVar)
	if err != nil {
		return Config{}, err
	}

	maxIdleConns, err := getEnvVarInt(maxIdleConnsEnvVar)
	if err != nil {
		return Config{}, err
	}

	maxConnsPerHost, err := getEnvVarInt(maxConnsPerHostEnvVar)
	if err != nil {
		return Config{}, err
	}

	idleConnTimeoutSecs, err := getEnvVarInt(idleConnTimeoutSecsEnvVar)
	if err != nil {
		return Config{}, err
	}

	dialerTimeoutSecs, err := getEnvVarInt(dialerTimeoutSecsEnvVar)
	if err != nil {
		return Config{}, err
	}

	dialerKeepAliveSecs, err := getEnvVarInt(dialerKeepAliveSecsEnvVar)
	if err != nil {
		return Config{}, err
	}

	tlsHandshakeTimeoutSecs, err := getEnvVarInt(tlsHandshakeTimeoutSecsEnvVar)
	if err != nil {
		return Config{}, err
	}

	disableKeepAlives, err := getEnvVarBool(disableKeepAlivesEnvVar)
	if err != nil {
		return Config{}, err
	}

	spaceXAPIEndpoint := os.Getenv(spaceXAPIEndpointEnvVar)
	if len(spaceXAPIEndpoint) == 0 {
		return Config{}, fmt.Errorf("unrecognised value for environment variable %s", spaceXAPIEndpointEnvVar)
	}

	cfg := Config{
		DBUsername:              username,
		DBPassword:              password,
		DBName:                  name,
		DBHost:                  host,
		DBMaxOpenConns:          openConns,
		DBMaxIdleConns:          idleConns,
		DBConnMaxLifetimeSecs:   connLifetime,
		DBConnRetries:           retries,
		DBConnRetryIntervalSecs: interval,
		APIPort:                 port,
		SwaggerPort:             swagPort,
		HTTPTimeout:             httpTimeout,
		MaxIdleConns:            maxIdleConns,
		MaxConnsPerHost:         maxConnsPerHost,
		IdleConnTimeoutSecs:     idleConnTimeoutSecs,
		DialerTimeoutSecs:       dialerTimeoutSecs,
		DialerKeepAliveSecs:     dialerKeepAliveSecs,
		TLSHandshakeTimeoutSecs: tlsHandshakeTimeoutSecs,
		DisableKeepAlives:       disableKeepAlives,
		SpaceXAPIEndpoint:       spaceXAPIEndpoint,
	}

	log.Println("Config loaded from environment variables")

	return cfg, nil
}

func getEnvVarInt(name string) (int, error) {
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

func getEnvVarBool(name string) (bool, error) {
	valueStr := os.Getenv(name)
	if len(valueStr) == 0 {
		return false, fmt.Errorf("unrecognised value for environment variable %s", name)
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return false, fmt.Errorf("can't parse environment variable %s", name)
	}

	return value, nil
}
