package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandlerOrders(t *testing.T) {
	// Create a multiplexer to run test on
	mux := mux.NewRouter()
	// Attaches handler you want to test
	mux.HandleFunc("/orders", handlerOrders)

	// Captures return HTTP response
	w := httptest.NewRecorder()
	// Creates request to handler you want to test
	r, _ := http.NewRequest("GET", "/orders", nil)

	// Send request to tested handler
	mux.ServeHTTP(w, r)

	// Check ResponseWriter for results
	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}
