package dynowebportal

import (
	"fmt"
	"net/http"
)

// RunWebPortal starts running the dino web portal on address addr
func RunWebPortal(addr string) (err error) {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(addr, nil)
	return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Dino web portal %s", r.RemoteAddr)
}
