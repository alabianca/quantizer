# Color Quantization
Reduce the number of colors in a source image. Still WIP and mostly exploring things for now.

## Usage
```go
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

## Example
[!alt text](https://github.com/alabianca/quantizer/blob/main/examples/basic/cobi.jpg)
