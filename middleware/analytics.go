package middleware

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"
)

var (
	ErrInvalidID    = errors.New("Invalid Id")
	ErrInvalidEmail = errors.New("Invalid Email")
)

func Recover(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				switch err {
				case ErrInvalidEmail:
					http.Error(w, ErrInvalidEmail.Error(), http.StatusUnauthorized)
				case ErrInvalidID:
					http.Error(w, ErrInvalidID.Error(), http.StatusUnauthorized)
				default:
					http.Error(w, "Unkwon error, recovered from panic", http.StatusInternalServerError)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// PassContext is used to pass values between middleware.
type PassContext func(ctx context.Context, w http.ResponseWriter, r *http.Request)

// ServeHTTP satisies the http.Handler interface.
func (fn PassContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(context.Background(), "foo", "bar")
	fn(ctx, w, r)
}
