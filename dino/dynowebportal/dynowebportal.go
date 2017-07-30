package dynowebportal

import "net/http"

// RunWebPortal starts running the dino web portal on address addr
func RunWebPortal(addr string) (err error) {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(addr, nil)
	return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

}
