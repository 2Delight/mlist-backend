package server

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	ResponseCounter.With(prometheus.Labels{
		operationNameLabel: r.RequestURI,
		statusCodeLabel:    strconv.Itoa(http.StatusOK),
	}).Inc()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
