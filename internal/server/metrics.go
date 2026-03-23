package server

import "github.com/prometheus/client_golang/prometheus"

const (
	operationNameLabel = "operation_name"
	statusCodeLabel    = "status_code"
)

var RequestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "mlist_request_total",
		Help: "RPS by operation counter",
	},
	[]string{operationNameLabel},
)

var ResponseCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "mlist_response_total",
		Help: "Responses by operation name and status code counter",
	},
	[]string{operationNameLabel, statusCodeLabel},
)

var Latency = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "mlist_latency_seconds",
		Buckets: []float64{
			0.001, 0.0025, 0.005, 0.0075,
			0.01, 0.025, 0.05, 0.075,
			0.1, 0.25, 0.5, 0.75,
			1, 2.5, 5, 7.5,
			10, 25, 50, 75,
		},
	},
	[]string{operationNameLabel},
)
