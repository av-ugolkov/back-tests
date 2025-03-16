package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestDurations = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "goobs_request_duration_seconds",
		Help:    "A histogram of the HTTP request durations in seconds.",
		Buckets: prometheus.DefBuckets,
	})

	RequestTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "goobs_request_total",
		Help: "A counter for the total number of requests received.",
	})

	Gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "goobs_gauge",
		Help: "A gauge for the current value of the gauge.",
	})
)

func New() *prometheus.Registry {
	registry := prometheus.NewRegistry()
	registry.MustRegister(RequestDurations, RequestTotal)

	return registry
}
