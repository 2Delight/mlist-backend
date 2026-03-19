package logger

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/log/global"
	sdk "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func SetupLogger(ctx context.Context, serviceName string, otlpTracesURL string, c Config) error {
	otlpLogsExporter, err := otlploggrpc.New(ctx, otlploggrpc.WithEndpointURL(otlpTracesURL))
	if err != nil {
		return err
	}

	stdoutLogsExporter, err := stdoutlog.New(stdoutlog.WithPrettyPrint())
	if err != nil {
		return err
	}

	resource, err := resource.New(
		ctx,
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)),
	)
	if err != nil {
		return err
	}

	loggerProvider := sdk.NewLoggerProvider(
		sdk.WithProcessor(sdk.NewBatchProcessor(stdoutLogsExporter)),
		sdk.WithProcessor(sdk.NewBatchProcessor(otlpLogsExporter)),
		sdk.WithResource(resource),
	)
	global.SetLoggerProvider(loggerProvider)

	err = setup(c)
	if err != nil {
		return err
	}

	return nil
}
