package server

import (
	"context"
	"net/http"
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

// func addLogging(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			logger.Info(r.Context(), "got request", "request", r)
// 			handler.ServeHTTP(w, r)
// 		},
// 	)
// }
