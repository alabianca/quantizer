package mediancut

import (
	"image"
	"image/color"
)

type Quantizer struct {
	Kind Algorithm
}

func (q *Quantizer) Quantize(src image.Image, colors color.Palette) (image.Image, error) {
	algo := q.Kind
	if algo == "" {
		algo = QuickSelect
	}

	return Quantize(src, colors, algo)
}
