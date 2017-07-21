package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Order struct {
	Id       int       `json:"id"`
	Products []Product `json:"product"`
	Client   string    `json:"client"`
}

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

var order Order

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", handlerOrders)

	http.ListenAndServe(":8000", router)
}

func handlerOrders(w http.ResponseWriter, r *http.Request) {
	orderRes, err := json.Marshal(createNewOrder())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(orderRes)
}

func handlerUpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func createNewOrder() Order {
	return Order{
		Id: 1,
		Products: []Product{
			Product{
				Id:    1,
				Name:  "Wallet",
				Price: 120.00,
			},
		},
		Client: "Miguel Angel",
	}
}
