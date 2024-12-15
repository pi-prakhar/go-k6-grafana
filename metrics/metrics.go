package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	totalRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
	)

	httpDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time for handler in seconds",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func InitMetrics() {
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(httpDuration)
}
