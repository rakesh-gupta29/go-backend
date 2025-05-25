package config

import (
	"time"

	"github.com/rakesh-gupta29/portal/internal/types"
)

var AppConfig types.ApplicationConfig

func loadAppConfig() {
	AppConfig.Host = "localhost"
	AppConfig.Port = "4000"
	AppConfig.DebugMode = true
	AppConfig.ReadTimeoutSecs = 10 * time.Second
}
