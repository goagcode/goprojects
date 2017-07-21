package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var router *mux.Router

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	router = mux.NewRouter()
}

func TestHandlerOrders(t *testing.T) {
	// Create a multiplexer to run test on
	// mux := mux.NewRouter()
	// Attaches handler you want to test
	router.HandleFunc("/orders", handlerOrders).Methods("GET")

	// Captures return HTTP response
	w := httptest.NewRecorder()
	// Creates request to handler you want to test
	r, _ := http.NewRequest("GET", "/orders", nil)

	// Send request to tested handler
	router.ServeHTTP(w, r)

	// Check ResponseWriter for results
	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}

	var order Order
	json.Unmarshal(w.Body.Bytes(), &order)
	if order.Id != 1 {
		t.Errorf("Cannot retrieve JSON order")
	}
}

func TestHandlerPutOrder(t *testing.T) {
	router.HandleFunc("/orders", handlerUpdateOrder).Methods("PUT")

	w := httptest.NewRecorder()
	order := strings.NewReader(`{
		"client": "Luis Angel"
	}`)
	r, _ := http.NewRequest("PUT", "/orders", order)

	router.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Status code expected 204 but got %v", w.Code)
	}
}
