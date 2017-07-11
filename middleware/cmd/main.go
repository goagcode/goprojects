package main

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/miguellgt/goprojects/middleware"
)

func panicker(w http.ResponseWriter, r *http.Request) {
	panic(middleware.ErrInvalidEmail)
}

func withContext(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	bar := ctx.Value("foo")
	w.Write((bar.([]byte)))
}

func main() {
	http.Handle("/panic", middleware.Recover(panicker))
	http.Handle("/context", middleware.PassContext(withContext))
	http.ListenAndServe(":3000", nil)
}
