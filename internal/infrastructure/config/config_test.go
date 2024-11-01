package config

import (
	"os"
	"reflect"
	"testing"
)

func Test_Get(t *testing.T) {
	const (
		user                       = "geoff"
		pwd                        = "password"
		name                       = "name"
		host                       = "localhost"
		dbOpenConnsStr             = "5"
		dbOpenConns                = 5
		dbIdleConnsStr             = "10"
		dbIdleConns                = 10
		dbConnLifetimeStr          = "1800"
		dbConnLifetime             = 1800
		dbRetriesStr               = "3"
		dbRetries                  = 3
		dbIntervalStr              = "2"
		dbInterval                 = 2
		port                       = ":3000"
		swagPort                   = ":3001"
		httpTimeoutStr             = "10"
		httpTimeout                = 10
		maxIdleConnsStr            = "5"
		maxIdleConns               = 5
		maxConnsPerHostStr         = "2"
		maxConnsPerHost            = 2
		idleConnTimeoutSecsStr     = "3"
		idleConnTimeoutSecs        = 3
		dialerTimeoutSecsStr       = "10"
		dialerTimeoutSecs          = 10
		dialerKeepAliveSecsStr     = "30"
		dialerKeepAliveSecs        = 30
		tlsHandshakeTimeoutSecsStr = "10"
		tlsHandshakeTimeoutSecs    = 10
		disableKeepAlivesStr       = "true"
		disableKeepAlives          = true
		spaceXAPIEndpoint          = "https://api.spacexdata.com"
	)

	tests := []struct {
		name                    string
		dbUserName              string
		dbPassword              string
		dbName                  string
		dbHost                  string
		dbOpenConns             string
		dbIdleConns             string
		dbConnLifeTime          string
		dbRetries               string
		dbInterval              string
		port                    string
		swagPort                string
		httpTimeout             string
		maxIdleConns            string
		maxConnsPerHost         string
		idleConnTimeoutSecs     string
		dialerTimeoutSecs       string
		dialerKeepAliveSecs     string
		tlsHandshakeTimeoutSecs string
		disableKeepAlives       string
		spaceXAPIEndpoint       string
		want                    Config
		wantErr                 string
	}{
		{
			name:                    "1. All env vars set, fully-populated Config returned",
			dbUserName:              user,
			dbPassword:              pwd,
			dbName:                  name,
			dbHost:                  host,
			dbOpenConns:             dbOpenConnsStr,
			dbIdleConns:             dbIdleConnsStr,
			dbConnLifeTime:          dbConnLifetimeStr,
			dbRetries:               dbRetriesStr,
			dbInterval:              dbIntervalStr,
			port:                    port,
			swagPort:                swagPort,
			httpTimeout:             httpTimeoutStr,
			maxIdleConns:            maxIdleConnsStr,
			maxConnsPerHost:         maxConnsPerHostStr,
			idleConnTimeoutSecs:     idleConnTimeoutSecsStr,
			dialerTimeoutSecs:       dialerTimeoutSecsStr,
			dialerKeepAliveSecs:     dialerKeepAliveSecsStr,
			tlsHandshakeTimeoutSecs: tlsHandshakeTimeoutSecsStr,
			disableKeepAlives:       disableKeepAlivesStr,
			spaceXAPIEndpoint:       spaceXAPIEndpoint,
			want: Config{
				DBUsername:              user,
				DBPassword:              pwd,
				DBName:                  name,
				DBHost:                  host,
				DBMaxOpenConns:          dbOpenConns,
				DBMaxIdleConns:          dbIdleConns,
				DBConnMaxLifetimeSecs:   dbConnLifetime,
				DBConnRetries:           dbRetries,
				DBConnRetryIntervalSecs: dbInterval,
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
			},
			wantErr: "",
		},
		{
			name:       "2. Missing env vars, empty Config and an error returned",
			dbUserName: "",
			dbPassword: pwd,
			dbName:     name,
			dbHost:     host,
			want:       Config{},
			wantErr:    "unrecognised value for environment variable DB_USERNAME",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				os.Unsetenv(dBUsernameEnvVar)
				os.Unsetenv(dbPasswordEnvVar)
				os.Unsetenv(dbNameEnvVar)
				os.Unsetenv(dbHostEnvVar)
				os.Unsetenv(dbMaxOpenConnsEnvVar)
				os.Unsetenv(dbMaxIdleConnsEnvVar)
				os.Unsetenv(dbConnMaxLifetimeSecsEnvVar)
				os.Unsetenv(dbConnRetriesEnvVar)
				os.Unsetenv(dbRetryIntervalEnvVar)
				os.Unsetenv(apiPortEnvVar)
				os.Unsetenv(swaggerPortEnvVar)
				os.Unsetenv(httpTimeoutEnvVar)
				os.Unsetenv(maxIdleConnsEnvVar)
				os.Unsetenv(maxConnsPerHostEnvVar)
				os.Unsetenv(idleConnTimeoutSecsEnvVar)
				os.Unsetenv(dialerTimeoutSecsEnvVar)
				os.Unsetenv(dialerKeepAliveSecsEnvVar)
				os.Unsetenv(tlsHandshakeTimeoutSecsEnvVar)
				os.Unsetenv(disableKeepAlivesEnvVar)
				os.Unsetenv(spaceXAPIEndpointEnvVar)

			}()

			if len(tt.dbUserName) > 0 {
				os.Setenv(dBUsernameEnvVar, tt.dbUserName)
			}
			if len(tt.dbPassword) > 0 {
				os.Setenv(dbPasswordEnvVar, tt.dbPassword)
			}
			if len(tt.dbName) > 0 {
				os.Setenv(dbNameEnvVar, tt.dbName)
			}
			if len(tt.dbHost) > 0 {
				os.Setenv(dbHostEnvVar, tt.dbHost)
			}
			if len(tt.dbOpenConns) > 0 {
				os.Setenv(dbMaxOpenConnsEnvVar, tt.dbOpenConns)
			}
			if len(tt.dbIdleConns) > 0 {
				os.Setenv(dbMaxIdleConnsEnvVar, tt.dbIdleConns)
			}
			if len(tt.dbConnLifeTime) > 0 {
				os.Setenv(dbConnMaxLifetimeSecsEnvVar, tt.dbConnLifeTime)
			}
			if len(tt.dbRetries) > 0 {
				os.Setenv(dbConnRetriesEnvVar, tt.dbRetries)
			}
			if len(tt.dbInterval) > 0 {
				os.Setenv(dbRetryIntervalEnvVar, tt.dbInterval)
			}
			if len(tt.port) > 0 {
				os.Setenv(apiPortEnvVar, tt.port)
			}
			if len(tt.swagPort) > 0 {
				os.Setenv(swaggerPortEnvVar, tt.swagPort)
			}
			if len(tt.httpTimeout) > 0 {
				os.Setenv(httpTimeoutEnvVar, tt.httpTimeout)
			}
			if len(tt.maxIdleConns) > 0 {
				os.Setenv(maxIdleConnsEnvVar, tt.maxIdleConns)
			}
			if len(tt.maxConnsPerHost) > 0 {
				os.Setenv(maxConnsPerHostEnvVar, tt.maxConnsPerHost)
			}
			if len(tt.idleConnTimeoutSecs) > 0 {
				os.Setenv(idleConnTimeoutSecsEnvVar, tt.idleConnTimeoutSecs)
			}
			if len(tt.dialerTimeoutSecs) > 0 {
				os.Setenv(dialerKeepAliveSecsEnvVar, tt.dialerKeepAliveSecs)
			}
			if len(tt.dialerKeepAliveSecs) > 0 {
				os.Setenv(dialerTimeoutSecsEnvVar, tt.dialerTimeoutSecs)
			}
			if len(tt.tlsHandshakeTimeoutSecs) > 0 {
				os.Setenv(tlsHandshakeTimeoutSecsEnvVar, tt.tlsHandshakeTimeoutSecs)
			}
			if len(tt.disableKeepAlives) > 0 {
				os.Setenv(disableKeepAlivesEnvVar, tt.disableKeepAlives)
			}
			if len(tt.spaceXAPIEndpoint) > 0 {
				os.Setenv(spaceXAPIEndpointEnvVar, tt.spaceXAPIEndpoint)
			}

			got, err := Get()

			if len(tt.wantErr) > 0 {
				if err == nil || (err.Error() != tt.wantErr) {
					t.Errorf("wrong error, got '%v', want '%s'", err, tt.wantErr)
					return
				}
			}
			if len(tt.wantErr) == 0 && err != nil {
				t.Errorf("wrong error, got '%v', want '%s'", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrong value, got = %v, want %v", got, tt.want)
			}
		})
	}
}
