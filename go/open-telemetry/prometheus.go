package main

import "github.com/prometheus/client_golang/prometheus"

var (
	requestCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "goapp_http_requests_total",
		Help: "Total number of HTTP requests handled by the Go app",
	})
	responseDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "goapp_response_duration_seconds",
		Help:    "Histogram of response durations for /hello",
		Buckets: prometheus.DefBuckets,
	})
)

func init() {
	prometheus.MustRegister(requestCount, responseDuration)
}
