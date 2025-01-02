package main

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/alabianca/quantizer"
)

func main() {
	// high def source image containing many colors
	reader, _ := os.Open("cobi.jpg")
	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	// want to reduce the source image to 256 colors
	colors := make([]quantizer.Point, 256)
	// imgCpy is the new image containing 256 colors
	imgCpy, err := quantizer.Quant(img, colors)
	if err != nil {
		panic(err)
	}

	out, _ := os.Create("output.jpeg")
	defer out.Close()

	if err := jpeg.Encode(out, imgCpy, nil); err != nil {
		panic(err)
	}

}
