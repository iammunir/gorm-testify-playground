package config

import (
	"os"
)

const (
	LOCAL       = "local"
	DEVELOPMENT = "development"
	PRODUCTION  = "production"
)

// ENVIRONMENT:
const ENVIRONMENT string = LOCAL // LOCAL, DEVELOPMENT, PRODUCTION

var env = map[string]map[string]string{
	// local environment configuration
	"local": {
		"PORT":                       "8260",
		"PQS_URL":                    "http://127.0.0.1:8765",
		"MYSQL_HOST":                 "localhost",
		"MYSQL_PORT":                 "3306",
		"MYSQL_USER":                 "root",
		"MYSQL_PASS":                 "admin",
		"MYSQL_SCHEMA":               "tutorials",
		"MONGO_URI":                  "mongodb://root:admin@localhost:27017",
		"MONGO_DATABASE":             "tutorial",
		"MONGO_POOL_MIN":             "10",
		"MONGO_POOL_MAX":             "100",
		"MONGO_MAX_IDLE_TIME_SECOND": "60",
	},

	// development environment configuration
	"development": {},

	// production environment configuration
	"production": {},
}

// CONFIG : global configuration
var CONFIG = env[ENVIRONMENT]

// Getenv : function for Environment Lookup
func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = Getenv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}
