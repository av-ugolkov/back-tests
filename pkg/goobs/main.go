package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"goobs/prom"
	"goobs/service"
	"goobs/tracer"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

var trHandler trace.Tracer
var svc *service.Service

func main() {
	ctx := context.Background()

	tr, err := tracer.InitTracer(ctx, "goobs-router")
	if err != nil {
		log.Fatalf("Error init Jaeger: %v", err)
	}
	defer func() { _ = tr.Shutdown(context.Background()) }()
	trHandler = tr.Tracer("goobs-tracer-handler")

	svc = service.New(trHandler)

	router := fiber.New()

	router.Get("/hello", handlerHello)
	router.Get("/hello/:name", handlerHello)
	router.Get("/metrics", adaptor.HTTPHandler(otelhttp.NewHandler(
		otelhttp.WithRouteTag("/metrics", http.HandlerFunc(metricsHandler)),
		"/metrics",
		otelhttp.WithTracerProvider(tr))))

	router.Listen(":8080")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	_, span := trHandler.Start(context.Background(), "metrics-handler")
	defer span.End()

	promhttp.HandlerFor(prom.New(),
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	).ServeHTTP(w, r)
}

func handlerHello(c *fiber.Ctx) error {
	ctx, span := trHandler.Start(c.Context(), "hello-handler")
	defer span.End()
	span.SetAttributes(attribute.String("time", time.Now().Format(time.DateTime)))

	prom.Gauge.Add(1)
	defer prom.Gauge.Sub(1)

	timer := prometheus.NewTimer(prom.RequestDurations)
	defer timer.ObserveDuration()

	span.AddEvent("start handler")
	defer span.AddEvent("finish handler")

	roll := 1 + rand.Intn(1000)
	time.Sleep(time.Duration(roll) * time.Millisecond)

	name := c.Params("name")
	if name == "" {
		name = "Unknown"
	}

	span.SetAttributes(attribute.String("params", name))

	prom.RequestTotal.Inc()

	propagator := propagation.TraceContext{}
	mp := make(map[string]string, 1)
	mp[name] = name
	propagator.Inject(ctx, propagation.MapCarrier(mp))

	svc.SomeAction(ctx, name, mp)

	msg := fmt.Sprintf("Hello, %s!", name)

	return c.SendString(msg)
}
