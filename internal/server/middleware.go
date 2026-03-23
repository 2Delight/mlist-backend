package server

import (
	"context"
	"net/http"
	"time"

	"github.com/2Delight/mlist-backend/internal/logger"
	"github.com/prometheus/client_golang/prometheus"
)

type middleware = func(http.Handler) http.Handler

func wrapHandlerFunc(handlerFunc http.HandlerFunc, middlewares ...middleware) http.Handler {
	var handler http.Handler = handlerFunc
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func addTimeout(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		},
	)
}

func addLogging(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info(r.Context(), "got request", "request", r.RequestURI)
			handler.ServeHTTP(w, r)
		},
	)
}

func addMetrics(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			RequestCounter.With(prometheus.Labels{
				operationNameLabel: r.RequestURI,
			}).Inc()
			startTime := time.Now()
			defer func() {
				finishTime := time.Now()
				Latency.With(prometheus.Labels{
					operationNameLabel: r.RequestURI,
				}).Observe(finishTime.Sub(startTime).Seconds())
			}()
			handler.ServeHTTP(w, r)
		},
	)
}
