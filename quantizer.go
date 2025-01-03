package quantizer

import (
	"image"
	"image/color"

	"github.com/alabianca/quantizer/mediancut"
)

type Quantizer interface {
	Quantize(src image.Image, colors color.Palette) (image.Image, error)
}

func Quantize(src image.Image, colors color.Palette) (image.Image, error) {
	im, err := mediancut.Quantize(src, colors, mediancut.QuickSelect)
	return im, err
}
