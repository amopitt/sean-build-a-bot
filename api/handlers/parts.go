package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/middleware"
)

func PartsFunction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		partsJSON, err := ioutil.ReadFile("./database/parts.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(partsJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleParts(mw *[]middleware.Middleware) http.Handler {
	// sign in
	partsHandler := http.HandlerFunc(PartsFunction)
	return middleware.WrapMiddleware(mw, partsHandler)
}
