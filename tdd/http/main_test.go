package main_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	. "github.com/miguellgt/goprojects/tdd/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakeOrder struct {
	Orders []Order
}

func (fo FakeOrder) Get() ([]Order, error) {
	return fo.Orders, nil
}

var _ = Describe("Get a post", func() {
	var router *mux.Router
	var w *httptest.ResponseRecorder
	var order *FakeOrder

	BeforeEach(func() {
		order = &FakeOrder{}
		// Create a multiplexer to run test on
		router = mux.NewRouter()
		// Attaches handler you want to test
		router.HandleFunc("/orders", HandlerOrders(&FakeOrder{})).Methods("GET")
		// Captures return HTTP response
		w = httptest.NewRecorder()
	})

	It("Get all orders", func() {
		// Creates request to handler you want to test
		r, _ := http.NewRequest("GET", "/orders", nil)
		// Send request to tested handler
		router.ServeHTTP(w, r)
		// Check ResponseWriter for results
		Expect(w.Code).To(Equal(200))
	})
})
