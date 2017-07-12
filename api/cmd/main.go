package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miguellgt/goprojects/api"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pages", api.CreatePage).Methods("POST")
	router.HandleFunc("/pages", api.GetPages).Methods("GET")
	router.HandleFunc("/pages/{pageId}", api.GetPage).Methods("GET")

	http.ListenAndServe(":3000", router)
}
