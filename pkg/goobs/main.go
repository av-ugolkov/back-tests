package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"goobs/prom"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := fiber.New()
	router.Get("/hello", handlerHello)
	router.Get("/hello/:name", handlerHello)
	router.Get("/metrics", adaptor.HTTPHandlerFunc(metricsHandler))

	router.Listen(":3000")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	promhttp.HandlerFor(prom.New(),
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	).ServeHTTP(w, r)
}

func handlerHello(c *fiber.Ctx) error {
	timer := prometheus.NewTimer(prom.RequestDurations)
	defer timer.ObserveDuration()

	roll := 1 + rand.Intn(1000)
	time.Sleep(time.Duration(roll) * time.Millisecond)

	name := c.Params("name")
	if name == "" {
		name = "Unknown"
	}

	prom.RequestTotal.Inc()
	msg := fmt.Sprintf("Hello, %s!", name)
	return c.SendString(msg)
}
