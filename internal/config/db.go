package config

import "github.com/rakesh-gupta29/portal/internal/types"

var DbConfig types.DatabaseConfig

func LoadDBConfig()  {
	DbConfig.Username  = "postgres"
	DbConfig.Password = "rakesh"
	DbConfig.Hostname = "localhost"
	DbConfig.Database = "portal_local"
}
