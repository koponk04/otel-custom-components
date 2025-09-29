package httpprocessor

import (
	"go.opentelemetry.io/collector/component"
)

// Config defines the configuration for the HTTP processor
type Config struct {
	// Endpoint is the HTTP endpoint to send data to
	Endpoint string `mapstructure:"endpoint"`
	
	// Headers are additional HTTP headers to include
	Headers map[string]string `mapstructure:"headers"`
	
	// Timeout for HTTP requests
	Timeout string `mapstructure:"timeout"`
}

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error {
	if cfg.Endpoint == "" {
		return component.ErrNilNextConsumer
	}
	return nil
}