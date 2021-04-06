package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/database"
	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/handlers"
	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/middleware"
	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/models"

	_ "github.com/lib/pq"
)

func main() {

	isProd := os.Getenv("PRODUCTION")
	if isProd != "" {
		database.InitProdEnvironment()
	}
	db, err := database.SetupDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	theStore := &models.Store{Db: db}

	fmt.Println("Successfully connected!")

	mw := &[]middleware.Middleware{middleware.CorsMiddleware, middleware.LoggingMiddleware}

	http.Handle("/api/sign-in", handlers.HandleSignIn(mw))
	http.Handle("/api/images/", http.StripPrefix("/api/images", handlers.HandleImages(mw)))
	http.Handle("/api/parts", handlers.HandleParts(mw))
	http.Handle("/api/cart", handlers.HandleCart(mw, theStore))

	// health check
	http.Handle("/api/health", handlers.HandleHealthCheck(mw))

	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	http.Handle("/api/ready", handlers.HandleReadyCheck(isReady, mw))

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Using default port 8085")
		port = "8085"
	}

	fmt.Printf("Listening on port %s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
