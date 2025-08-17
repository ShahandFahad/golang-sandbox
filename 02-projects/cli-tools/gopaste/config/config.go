package config

// Package config centralizes configuration loading and validation.
// We prefer environment-driven config (12-factor), with .env for local dev.

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config holds all runtime configuration the app needs.
// Keep fields simple and explicit; avoid "magical" implicit defaults.
type Config struct {
	APIBaseURL  string        // e.g. https://api.my-paste.example
	APIKey      string        // optional; set if provider requires auth
	HTTPTimeout time.Duration // request timeout; defaults to 10s
}

// App is the global runtime configuration once loaded.
// In larger apps we'd inject Config instead of using globals, but for
// a CLI this is pragmatic and keeps command wiring simple.
var App Config

// Load reads configuration from .env (if present) and environment variables.
// Precedence: env vars > .env > defaults
func Load() error {
	v := viper.New()

	// Allow local development with a .env file; not required in CI/prod.
	v.SetConfigFile(".env")
	v.SetConfigType("env")

	// Read environment variables automatically as overrides.
	// Example: export API_BASE_URL=https://...
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Sensible defaults for a CLI tool.
	v.SetDefault("API_BASE_URL", "")
	v.SetDefault("API_KEY", "")
	v.SetDefault("HTTP_TIMEOUT", "10s")

	// .env is optional—warn but don’t fail if missing.
	if err := v.ReadInConfig(); err != nil {
		log.Printf("config: no .env file found: %v (using env/defaults)", err)
	}

	App.APIBaseURL = v.GetString("API_BASE_URL")
	App.APIKey = v.GetString("API_KEY")

	// viper parses Go duration strings like "10s", "2m".
	App.HTTPTimeout = v.GetDuration("HTTP_TIMEOUT")
	if App.HTTPTimeout <= 0 {
		App.HTTPTimeout = 10 * time.Second
	}

	return nil
}
