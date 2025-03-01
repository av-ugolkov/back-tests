package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	jaegerEndpoint = "localhost:4317"
	serviceName    = "Fibonacci"
)

func newTracerProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to merge resources: %w", err)
	}

	stdExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build StdoutExporter: %w", err)
	}

	otlpExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(jaegerEndpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build OtlpExporter: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(stdExporter),
		sdktrace.WithBatcher(otlpExporter),
	)

	return tp, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	tp, err := newTracerProvider(ctx)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return
	}
	defer func() { tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(fibHandler), "root"))

	if err := http.ListenAndServe(":3000", nil); err != nil {
		slog.ErrorContext(ctx, err.Error())
		return
	}
}

func fibHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	sp := trace.SpanFromContext(ctx)
	args := req.URL.Query()["n"]
	if len(args) != 1 {
		msg := "wrong number of arguments"
		sp.SetStatus(codes.Error, msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	sp.SetAttributes(attribute.String("fibonacci.argument", args[0]))

	n, err := strconv.Atoi(args[0])
	if err != nil {
		msg := fmt.Sprintf("couldn't parse index n: %s", err.Error())
		sp.SetStatus(codes.Error, msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	sp.SetAttributes(attribute.Int("fibonacci.parameter", n))

	result := Fibonacci(ctx, n)
	sp.SetAttributes(attribute.Int("fibonacci.result", result))
	fmt.Fprintln(w, result)
}

func Fibonacci(ctx context.Context, n int) int {
	ctx, sp := otel.GetTracerProvider().Tracer(serviceName).Start(
		ctx,
		"Fibonacci",
		trace.WithAttributes(attribute.Int("fibonacci.n", n)),
	)
	defer sp.End()

	result := 1
	if n > 1 {
		a := Fibonacci(ctx, n-1)
		b := Fibonacci(ctx, n-2)
		result = a + b
	}

	sp.SetAttributes(attribute.Int("fibonacci.result", result))

	return result
}
