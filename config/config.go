package config

import (
	"net/http"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/confmap"
)

var Config = koanf.New(".")

func Defaults(c *koanf.Koanf) error {
		return c.Load(confmap.Provider(map[string]interface{} {
			// Logger Defaults
			"logger.level": 			 "info",
			"logger.encoding": 			 "console",
			"logger.color": 			 true,
			"logger.dev_mode":           true,
			"logger.disable_caller":     false,
			"logger.disable_stacktrace": true,

			// Server Configuration
			"server.host": "",
			"server.port": "8080",
			"server.cors.allowed_origins":     []string{"*"},
			"server.cors.allowed_methods":     []string{http.MethodHead, http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
			"server.cors.allowed_headers":     []string{"*"},
			"server.cors.allowed_credentials": false,
			"server.cors.max_age":             300,

			// Database Settings
			"database.username":              "postgres",
			"database.password":              "password",
			"database.host":                  "postgres",
			"database.port":                  5432,
			"database.database":              "gorestapi",
			"database.auto_create":           true,
			"database.search_path":           "",
			"database.sslmode":               "disable",
			"database.sslcert":               "",
			"database.sslkey":                "",
			"database.sslrootcert":           "",
			"database.retries":               5,
			"database.sleep_between_retries": "7s",
			"database.max_connections":       40,
			"database.log_queries":           false,
			"database.wipe_confirm":          false,
		}, "."), nil)
}