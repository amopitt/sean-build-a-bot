package handlers

import (
	"net/http"

	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/middleware"
)

// HealthCheck is a liveness probe.
func HealthCheckFunction(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandleHealthCheck(mw *[]middleware.Middleware) http.Handler {
	// sign in
	healthCheckHandler := http.HandlerFunc(HealthCheckFunction)
	// commenting out middleware to not log repeatedly
	return healthCheckHandler // middleware.WrapMiddleware(mw, healthCheckHandler)
}
