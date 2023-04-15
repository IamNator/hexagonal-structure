package api

import (
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()

    // register order routes
    r.HandleFunc("/orders/{id}", GetOrderHandler).Methods("GET")
    r.HandleFunc("/orders", CreateOrderHandler).Methods("POST")

    // apply middleware
    r.Use(LoggingMiddleware)

    return r
}
