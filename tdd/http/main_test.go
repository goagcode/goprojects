package main_test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
	. "github.com/miguellgt/goprojects/tdd/http"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

type FakeOrder struct {
	Orders []Order
}

func (fo FakeOrder) Get() ([]Order, error) {
	return fo.Orders, nil
}

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

	It("Put Order", func() {
		router := mux.NewRouter()
		router.HandleFunc("/orders", HandlerUpdateOrder).Methods("PUT")
		w := httptest.NewRecorder()
		order := strings.NewReader(`{
			"client": "Luis Angel"
		}`)
		r, _ := http.NewRequest("PUT", "/orders", order)

		router.ServeHTTP(w, r)

		if w.Code != 204 {
			GinkgoT().Errorf("Status code expected 204 but got %v", w.Code)
		}
	})
})
