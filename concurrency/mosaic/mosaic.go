package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"os"
)

func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1 := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

func rezise(in image.Image, newWith int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx() / newWith
	out := image.NewRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.X/ratio, bounds.Max.X/ratio, bounds.Max.Y/ratio))
	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+radio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i := x+radio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return *out
}

func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles DB ...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles")
	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILEDB: ", err, name)
			}
		} else {
			fmt.Println("cannot open file", name, err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles DB.")
	return db
}

func main() {
	fmt.Println("mosaic program")
}
