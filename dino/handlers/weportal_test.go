package handlers_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/miguellgt/goprojects/dino/handlers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Weportal", func() {
	Describe("When client send request to GET / endpoint", func() {
		It("should return a message 'Welcome to Dino Web Portal'", func() {
			mux := mux.NewRouter()
			mux.HandleFunc("/", handlers.RootHandler).Methods("GET")

			w := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/", nil)
			mux.ServeHTTP(w, r)

			Expect(w.Code).To(Equal(200))
			Expect(string(w.Body.Bytes())).To(Equal("Welcome to Dino Web Portal"))
		})
	})
})
