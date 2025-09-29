package httpprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

const (
	typeStr   = "httpprocessor"
	stability = component.StabilityLevelAlpha
)

// NewFactory creates a new processor factory
func NewFactory() processor.Factory {
	return processor.NewFactory(
		component.MustNewType(typeStr),
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, stability),
		processor.WithMetrics(createMetricsProcessor, stability),
		processor.WithLogs(createLogsProcessor, stability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		// Add your default configuration here
		Endpoint: "http://localhost:8080",
	}
}

func createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	processorCfg := cfg.(*Config)
	proc := &httpProcessor{
		config: processorCfg,
		logger: set.Logger,
		next:   nextConsumer,
	}

	return proc, nil
}

func createMetricsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {
	processorCfg := cfg.(*Config)
	proc := &httpMetricsProcessor{
		config: processorCfg,
		logger: set.Logger,
		next:   nextConsumer,
	}

	return proc, nil
}

func createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	processorCfg := cfg.(*Config)
	proc := &httpLogsProcessor{
		config: processorCfg,
		logger: set.Logger,
		next:   nextConsumer,
	}

	return proc, nil
}