package main

import (
	"net/http"

	"github.com/miguellgt/goprojects/middleware"
)

func panicker(w http.ResponseWriter, r *http.Request) {
	panic("Wahhhhh")
}

func main() {
	// logger := middleware.CreateLogger("section4")
	// http.Handle("/", middleware.Time(logger, hello))
	http.Handle("/panic", middleware.Recover(panicker))
	http.ListenAndServe(":3000", nil)
}
