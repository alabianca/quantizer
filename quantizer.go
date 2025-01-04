package quantizer

import (
	"image"
	"image/color"

	"github.com/alabianca/quantizer/mediancut"
)

type Quantizer interface {
	Quantize(src image.Image, colors color.Palette) (image.Image, error)
}

// Quantize uses the default image quantization algorithm "median cut"
// to quantize src using the provided colors palette. The quantized image
// will be returned and will have the same dimensions as the provided src image
func Quantize(src image.Image, colors color.Palette) (image.Image, error) {
	im, err := mediancut.Quantize(src, colors, mediancut.QuickSelect)
	return im, err
}
