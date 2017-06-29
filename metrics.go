package main

import "github.com/prometheus/client_golang/prometheus"

var (
	httpRequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   cfg.Namespace,
		Name:        "http_response_count_total",
		Help:        "Amount of processed HTTP requests",
		ConstLabels: prometheus.Labels{},
	}, []string{"code", "method", "endpoint"})
	httpResponseBodySizeBytes = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   cfg.Namespace,
		Name:        "http_response_size_bytes",
		Help:        "Total amount of transferred bytes",
		ConstLabels: prometheus.Labels{},
	}, []string{"code", "method", "endpoint"})
)

func init() {
	prometheus.MustRegister(httpRequestCount)
	prometheus.MustRegister(httpResponseBodySizeBytes)
}
