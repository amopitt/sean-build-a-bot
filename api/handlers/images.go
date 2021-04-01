package handlers

import (
	"net/http"

	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/middleware"
)

func HandleImages(mw []middleware.Middleware) http.Handler {
	fileServer := http.FileServer((http.Dir("./images")))
	return middleware.WrapMiddleware(mw, fileServer)
}
