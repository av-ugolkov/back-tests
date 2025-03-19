package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"goobs/prom"
	"goobs/tracer"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var otracer trace.Tracer

func main() {
	ctx := context.Background()

	tp, err := tracer.InitTracer(ctx)
	if err != nil {
		log.Fatalf("Error init Jaeger: %v", err)
	}
	defer func() { _ = tp.Shutdown(context.Background()) }()
	otracer = otel.Tracer("goobs-tracer")

	router := fiber.New()
	router.Get("/hello", handlerHello)
	router.Get("/hello/:name", handlerHello)
	router.Get("/metrics", adaptor.HTTPHandlerFunc(metricsHandler))

	router.Listen(":8080")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	_, span := otracer.Start(context.Background(), "metrics-handler")
	defer span.End()

	promhttp.HandlerFor(prom.New(),
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	).ServeHTTP(w, r)
}

func handlerHello(c *fiber.Ctx) error {
	prom.Gauge.Add(1)
	defer prom.Gauge.Sub(1)

	timer := prometheus.NewTimer(prom.RequestDurations)
	defer timer.ObserveDuration()

	roll := 1 + rand.Intn(1000)
	time.Sleep(time.Duration(roll) * time.Millisecond)

	name := c.Params("name")
	if name == "" {
		name = "Unknown"
	}

	_, span := otracer.Start(context.Background(), "hello-handler")
	span.SetAttributes(attribute.String("params", name))
	defer span.End()

	prom.RequestTotal.Inc()
	msg := fmt.Sprintf("Hello, %s!", name)
	return c.SendString(msg)
}
