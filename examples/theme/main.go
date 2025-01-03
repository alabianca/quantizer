package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/alabianca/quantizer/mediancut"
)

func main() {
	reader, err := os.Open("./cobi.jpg")
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	img, err := jpeg.Decode(reader)
	if err != nil {
		panic(err)
	}

	colors := make(color.Palette, 8)
	_, err = mediancut.Quantize(img, colors, mediancut.QuickSelect)
	if err != nil {
		panic(err)
	}

	// Create a new image with a white background
	im := image.NewRGBA(image.Rect(0, 0, 200, 200))
	white := color.RGBA{255, 255, 255, 255}
	for x := 0; x < 200; x++ {
		for y := 0; y < 200; y++ {
			im.Set(x, y, white)
		}
	}

	// Draw the theme
	drawTheme(colors, im, 200, 200)

	// Save the image as a PNG file
	f, _ := os.Create("theme_quickselect.png")
	defer f.Close()
	png.Encode(f, im)

	fmt.Println("Done!")
}

func drawTheme(colors color.Palette, img *image.RGBA, xMax, yMax int) {
	for i, c := range colors {
		//c := color.RGBA{uint8(c.Red), uint8(c.Green), uint8(c.Blue), 255}
		yStart := (yMax / len(colors)) * i
		for x := 0; x < xMax; x++ {
			for y := yStart; y < yMax; y++ {
				img.Set(x, y, c)
			}
		}
	}
}
