package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/middleware"
	"github.com/amopitt/sean-build-a-bot/api/build-a-bot-api/models"
)

func CartFunction(w http.ResponseWriter, r *http.Request, store *models.Store) {
	switch r.Method {
	case http.MethodPost:

		decoder := json.NewDecoder(r.Body)
		var cart models.Cart
		err := decoder.Decode(&cart)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		fmt.Println(cart.RobotName)
		sqlStatement := `
		INSERT INTO orders (robot_name, cost, created_on) 
		VALUES ($1, $2, $3)
		RETURNING order_id
		`
		id := 0
		err = store.Db.QueryRow(sqlStatement, cart.RobotName, cart.Cost, time.Now()).Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("New record ID is:", id)
		w.WriteHeader(http.StatusCreated)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleCart(mw *[]middleware.Middleware, store *models.Store) http.Handler {
	//cartHandler := http.HandlerFunc(CartFunction)
	cartHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CartFunction(w, r, store)
	})
	return middleware.WrapMiddleware(mw, cartHandler)
}
