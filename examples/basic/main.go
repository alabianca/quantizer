package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/alabianca/quantizer/mediancut"
)

func main() {
	// high def source image containing many colors
	reader, _ := os.Open("cobi.jpg")
	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	// want to reduce the source image to 256 colors
	colors := make(color.Palette, 8)
	// imgCpy is the new image containing 256 colors
	imgCpy, err := mediancut.Quantize(img, colors, mediancut.QuickSelect)
	if err != nil {
		panic(err)
	}

	out, _ := os.Create("output_quickselect.jpeg")
	defer out.Close()

	if err := jpeg.Encode(out, imgCpy, nil); err != nil {
		panic(err)
	}

}
