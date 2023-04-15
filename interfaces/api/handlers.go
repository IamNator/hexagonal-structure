package api

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
    // retrieve order data from the database
    order, err := db.GetOrderByID(id)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error retrieving order with ID %d: %s", id, err.Error()), http.StatusInternalServerError)
        return
    }

    // encode order data as JSON and write to the response
    json.NewEncoder(w).Encode(order)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
    // decode order data from request body
    var order domain.Order
    err := json.NewDecoder(r.Body).Decode(&order)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error decoding order data: %s", err.Error()), http.StatusBadRequest)
        return
    }

    // create order in the database
    err = db.CreateOrder(&order)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error creating order: %s", err.Error()), http.StatusInternalServerError)
        return
    }

    // encode order data as JSON and write to the response
    json.NewEncoder(w).Encode(order)
}
