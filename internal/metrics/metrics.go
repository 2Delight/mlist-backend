package metrics

import (
	"fmt"
	"net/http"

	"github.com/2Delight/mlist-backend/internal/server"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupMetrics(port uint16) {
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		// collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		server.RequestCounter,
		server.ResponseCounter,
		server.Latency,
	)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
