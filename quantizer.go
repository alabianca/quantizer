package quantizer

import (
	"image"
	"image/color"

	"github.com/alabianca/quantizer/mediancut"
)

// type Quantizer interface {
// 	Quantize(src image.Image, colors []Point) (image.Image, error)
// }

func Quantize(src image.Image, colors []color.RGBA64) (image.Image, error) {
	im, err := mediancut.Quantize(src, make([]mediancut.Point, len(colors)), mediancut.MergeSort)
	return im, err
}
