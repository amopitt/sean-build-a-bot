package handlers

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/middleware"
)

// ReadyCheck is a readiness probe.
func ReadyCheckFunction(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func HandleReadyCheck(isReady *atomic.Value, mw *[]middleware.Middleware) http.Handler {
	// ready check
	readyCheckFunction := func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Ready spaghetti!")
	}

	readyCheckHandler := http.HandlerFunc(readyCheckFunction)
	// commenting out middleware to not log repeatedly
	return readyCheckHandler // #middleware.WrapMiddleware(mw, readyCheckHandler)
}
