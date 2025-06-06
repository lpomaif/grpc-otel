package gotel

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

// Config holds the configuration for the telemetry.
type Config struct {
	ServiceName    string        `env:"TELEMETRY_SERVICE_NAME"      envDefault:"gotel"`
	ServiceVersion string        `env:"TELEMETRY_SERVICE_VERSION"   envDefault:"0.0.1"`
	Environment    string        `env:"TELEMETRY_ENVIRONMENT"       envDefault:"development"`
	Enabled        bool          `env:"TELEMETRY_ENABLED" envDefault:"true"`
	GrpcUrl        string        `env:"TELEMETRY_GRPC_URL"          envDefault:"localhost:4317"`
	Timeout        time.Duration `env:"TELEMETRY_TIMEOUT"           envDefault:"5s"`
}

// NewConfigFromEnv creates a new telemetry config from the environment.
func NewConfigFromEnv() (*Config, error) {
	telem := &Config{}
	if err := env.Parse(telem); err != nil {
		return &Config{}, fmt.Errorf("failed to parse telemetry config: %w", err)
	}

	return telem, nil
}
