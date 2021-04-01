package middleware

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request: " + r.URL.Path)
		handler.ServeHTTP(w, r)
		fmt.Println("Completed request: " + r.URL.Path)
	})
}
