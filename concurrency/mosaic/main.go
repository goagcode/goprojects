package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
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

	newImage := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))
	// Clones tile database
	db := cloneTilesDB()
	// Set up source point for each tile
	sp := image.Point{0, 0}
	// Iterates through target image
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.x; x < bounds.Max.X; x + tileSize {
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}
			nearest := nearest(color, &db)
			file, err := os.Open(nearest)
			if err == nil {
				img, _, err := image.Decode(file)
				if err == nil {
					t := resize(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.React(x, y, x+tileSize, y+tileSize)
					draw.Draw(newImage, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println("error: ", err, nearest)
				}
			} else {
				fmt.Println("error: ", nearest)
			}
			file.Close()
		}
	}
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
