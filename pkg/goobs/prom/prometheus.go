package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestDurations = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "goobs_request_duration_seconds",
		Help: "A histogram of the HTTP request durations in seconds.",
	})
)

func New() *prometheus.Registry {
	registry := prometheus.NewRegistry()
	registry.MustRegister(requestDurations)

	return registry
}
