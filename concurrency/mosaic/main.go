package main

import (
	"fmt"
	"html/template"
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

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}
