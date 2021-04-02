package handlers

import (
	"net/http"
	"time"

	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/middleware"
)

func CartFunction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		time.Sleep(800 * time.Millisecond)
		w.WriteHeader(http.StatusCreated)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleCart(mw *[]middleware.Middleware) http.Handler {
	cartHandler := http.HandlerFunc(CartFunction)
	return middleware.WrapMiddleware(mw, cartHandler)
}
