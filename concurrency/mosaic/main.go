package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"
	"time"
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

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	r.ParseMultipartForm(10485760)
	// Gets upload file and tile size
	file, _ := r.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))
	// Decodes uploaded target image
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	// Clones tile database
	db := cloneTilesDB()
	// Fanning out,cutting up image for independent processing
	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := cut(original, &db, tileSize, bounds.Min.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, tileSize, bounds.Min.X/2, bounds.Min.Y/2, bounds.Max.X, bounds.Max.Y)
	c := combine(bounds, c1, c2, c3, c4)
	// Encoding in JPEG, deliver to browser in base64 string
	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newImage, nil)
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())
	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":   mosaic,
		"duration": fmt.Sprintf("%v ", t1.Sub(0)),
	}
	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)
}
