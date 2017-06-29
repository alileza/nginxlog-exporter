package main

import "github.com/prometheus/client_golang/prometheus"

var (
	httpRequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:        "http_requests_total",
		Help:        "Number status 2xx",
		ConstLabels: prometheus.Labels{},
	}, []string{"code", "method", "endpoint"})
	httpResponseBodySizeBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:        "http_response_body_size_bytes",
		Help:        "Response body size",
		ConstLabels: prometheus.Labels{},
	}, []string{"code", "method", "endpoint"})
)

func init() {
	prometheus.MustRegister(httpRequestCount)
	prometheus.MustRegister(httpResponseBodySizeBytes)
}
