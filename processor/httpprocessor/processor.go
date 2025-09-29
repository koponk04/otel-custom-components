package httpprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

// httpProcessor handles traces
type httpProcessor struct {
	config *Config
	logger *zap.Logger
	next   consumer.Traces
}

func (p *httpProcessor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}

func (p *httpProcessor) Start(ctx context.Context, host component.Host) error {
	return nil
}

func (p *httpProcessor) Shutdown(ctx context.Context) error {
	return nil
}

func (p *httpProcessor) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	// Implement your trace processing logic here
	p.logger.Info("Processing traces", zap.Int("span_count", td.SpanCount()))
	
	// Example: Add an attribute to all spans
	resourceSpans := td.ResourceSpans()
	for i := 0; i < resourceSpans.Len(); i++ {
		rs := resourceSpans.At(i)
		scopeSpans := rs.ScopeSpans()
		for j := 0; j < scopeSpans.Len(); j++ {
			ss := scopeSpans.At(j)
			spans := ss.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				span.Attributes().PutStr("custom.processor", "httpprocessor")
			}
		}
	}
	
	return p.next.ConsumeTraces(ctx, td)
}

// httpMetricsProcessor handles metrics
type httpMetricsProcessor struct {
	config *Config
	logger *zap.Logger
	next   consumer.Metrics
}

func (p *httpMetricsProcessor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}

func (p *httpMetricsProcessor) Start(ctx context.Context, host component.Host) error {
	return nil
}

func (p *httpMetricsProcessor) Shutdown(ctx context.Context) error {
	return nil
}

func (p *httpMetricsProcessor) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	// Implement your metric processing logic here
	p.logger.Info("Processing metrics", zap.Int("metric_count", md.MetricCount()))
	
	return p.next.ConsumeMetrics(ctx, md)
}

// httpLogsProcessor handles logs
type httpLogsProcessor struct {
	config *Config
	logger *zap.Logger
	next   consumer.Logs
}

func (p *httpLogsProcessor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}

func (p *httpLogsProcessor) Start(ctx context.Context, host component.Host) error {
	return nil
}

func (p *httpLogsProcessor) Shutdown(ctx context.Context) error {
	return nil
}

func (p *httpLogsProcessor) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	// Implement your log processing logic here
	p.logger.Info("Processing logs", zap.Int("log_count", ld.LogRecordCount()))
	
	// Example: Add an attribute to all log records
	resourceLogs := ld.ResourceLogs()
	for i := 0; i < resourceLogs.Len(); i++ {
		rl := resourceLogs.At(i)
		scopeLogs := rl.ScopeLogs()
		for j := 0; j < scopeLogs.Len(); j++ {
			sl := scopeLogs.At(j)
			logRecords := sl.LogRecords()
			for k := 0; k < logRecords.Len(); k++ {
				lr := logRecords.At(k)
				lr.Attributes().PutStr("custom.processor", "httpprocessor")
			}
		}
	}
	
	return p.next.ConsumeLogs(ctx, ld)
}