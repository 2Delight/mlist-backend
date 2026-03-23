package server

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	const operationName = "ping"
	RequestCounter.With(prometheus.Labels{
		operationNameLabel: operationName,
	}).Inc()
	ResponseCounter.With(prometheus.Labels{
		operationNameLabel: operationName,
		statusCodeLabel:    strconv.Itoa(http.StatusOK),
	}).Inc()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
