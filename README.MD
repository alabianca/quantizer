# Image Quantizer
Reduce the number of colors in a source image. Still WIP and mostly exploring things for now.

## Basic Usage
```go
func main() {
	// high def source image containing many colors
	reader, _ := os.Open("cobi.jpg")
	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	// want to reduce the source image to 256 colors
	colors := make(color.Palette, 256)
	// quantize img using the default algorithm median cut
	imgCpy, err := quantizer.Quantize(img, colors)
	if err != nil {
		panic(err)
	}

	out, _ := os.Create("output.jpeg")
	defer out.Close()

	if err := jpeg.Encode(out, imgCpy, nil); err != nil {
		panic(err)
	}

}
```

### Result
<img src="https://github.com/alabianca/quantizer/blob/main/docs/example_quantization.png"/>

### Theme usage
It's somewhat possible to generate a theme/color palette that fits fairly well with the provided source image
```go
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

	// quantize img into 8 colors. This will be our theme
	colors := make(color.Palette, 8)
	_, err = quantizer.Quantize(img, colors)
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

	// Draw the theme using the colors that were
	// extracted from the src image onto a 200x200 img
	drawTheme(colors, im, 200, 200)

	// Save the image as a PNG file
	f, _ := os.Create("theme.png")
	defer f.Close()
	png.Encode(f, im)

	fmt.Println("Done!")
}

func drawTheme(colors color.Palette, img *image.RGBA, xMax, yMax int) {
	for i, c := range colors {
		yStart := (yMax / len(colors)) * i
		for x := 0; x < xMax; x++ {
			for y := yStart; y < yMax; y++ {
				img.Set(x, y, c)
			}
		}
	}
}
```

### Result
<img src="https://github.com/alabianca/quantizer/blob/main/examples/theme/theme.png"/>