package server

import "github.com/prometheus/client_golang/prometheus"

const (
	operationNameLabel = "operation_name"
	statusCodeLabel    = "status_code"
)

var RequestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "mlist_request_counter",
		Help: "RPS by operation counter",
	},
	[]string{operationNameLabel},
)

var ResponseCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "mlist_response_counter",
		Help: "Responses by operation name and status code counter",
	},
	[]string{operationNameLabel, statusCodeLabel},
)
