package handlers

import (
	"net/http"

	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/middleware"
)

func SignInFunction(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandleSignIn(mw *[]middleware.Middleware) http.Handler {
	// sign in
	signInHandler := http.HandlerFunc(SignInFunction)
	return middleware.WrapMiddleware(mw, signInHandler)
}
