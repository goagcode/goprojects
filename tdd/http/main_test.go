package main_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	. "github.com/miguellgt/goprojects/tdd/http"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Testing Order Resource", func() {
	It("Get orders", func() {
		// Create a multiplexer to run test on
		router := mux.NewRouter()
		// Captures return HTTP response
		w := httptest.NewRecorder()
		// Attaches handler you want to test
		router.HandleFunc("/orders", HandlerOrders(&FakeOrder{})).Methods("GET")
		// Creates request to handler you want to test
		r, _ := http.NewRequest("GET", "/orders", nil)
		// Send request to tested handler
		router.ServeHTTP(w, r)

		// Check ResponseWriter for results
		if w.Code != 200 {
			GinkgoT().Errorf("Response code is %v", w.Code)
		}
	})
})
