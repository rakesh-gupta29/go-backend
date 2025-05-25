package types

import (
	"time"
)

type DatabaseConfig struct {
	Username string
	Password string
	Hostname string
	Database string
}

type ApplicationConfig struct {
	Host            string
	Port            string
	DebugMode       bool
	ReadTimeoutSecs time.Duration
}
