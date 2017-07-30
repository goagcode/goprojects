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

	Describe("RootHandler", func() {
		var router *mux.Router
		var w *httptest.ResponseRecorder

		BeforeEach(func() {
			router = mux.NewRouter()
			router.HandleFunc("/", handlers.RootHandler).Methods("GET")
			w = httptest.NewRecorder()
		})

		Context("When client send a request to GET / endpoint", func() {
			BeforeEach(func() {
				r, _ := http.NewRequest("GET", "/", nil)
				router.ServeHTTP(w, r)
			})

			It("should return a message 'Welcome to Dino Web Portal'", func() {
				Expect(w.Code).To(Equal(200))
				Expect(string(w.Body.Bytes())).To(Equal("Welcome to Dino Web Portal"))
			})

			It("should return 200 statuc code", func() {
				Expect(w.Code).To(Equal(200))
			})
		})
	})
})
