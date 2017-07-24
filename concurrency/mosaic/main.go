package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)

	server := &http.Server{
		Addr:    "8080",
		Handler: mux,
	}

	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.")
	server.ListenAndServe()
}
