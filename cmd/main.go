package main

import (
	"context"
	"fmt"
	"grpc-otel/pkg/gotel"
	"os"
)

func main() {
	ctx := context.Background()

	config, err := gotel.NewConfigFromEnv()
	if err != nil {
		fmt.Println("failed to load telemetry config")
		os.Exit(1)
	}

	// Initialize telemetry. If the exporter fails, fallback to nop.
	var telem gotel.TelemetryProvider
	telem, err = gotel.NewTelemetry(ctx, config)
	if err != nil {
		fmt.Println("failed to create telemetry, falling back to no-op telemetry")
		telem, _ = gotel.NewNoopTelemetry(config)
	}
	defer telem.Shutdown(ctx)

	telem.LogInfo("telemetry initialized")
}
