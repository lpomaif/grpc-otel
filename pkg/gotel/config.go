package gotel

import (
	"fmt"

	"github.com/caarlos0/env"
)

// Config holds the configuration for the telemetry.
type Config struct {
	ServiceName    string `env:"TELEMETRY_SERVICE_NAME"      envDefault:"gotel"`
	ServiceVersion string `env:"TELEMETRY_SERVICE_VERSION"   envDefault:"0.0.1"`
	Environment    string `env:"TELEMETRY_ENVIRONMENT"       envDefault:"development"`
	Enabled        bool   `env:"TELEMETRY_ENABLED" envDefault:"true"`
	GrpcUrl        string `env:"TELEMETRY_GRPC_URL"          envDefault:"localhost:4317"`
}

// NewConfigFromEnv creates a new telemetry config from the environment.
func NewConfigFromEnv() (Config, error) {
	telem := Config{}
	if err := env.Parse(&telem); err != nil {
		return Config{}, fmt.Errorf("failed to parse telemetry config: %w", err)
	}

	return telem, nil
}
