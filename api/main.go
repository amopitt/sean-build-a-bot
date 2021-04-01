package main

import (
	"fmt"
	"net/http"

	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/handlers"
	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/middleware"
)

const port = "8085"

func main() {

	mw := []middleware.Middleware{middleware.CorsMiddleware, middleware.LoggingMiddleware}

	http.Handle("/api/sign-in", handlers.HandleSignIn(mw))
	http.Handle("/api/images/", http.StripPrefix("/api/images", handlers.HandleImages(mw)))
	http.Handle("/api/parts", handlers.HandleParts(mw))
	http.Handle("/api/cart", handlers.HandleCart(mw))

	fmt.Printf("Listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
