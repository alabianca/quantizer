package mediancut

import (
	"image"
	"image/color"
	"math"
)

type Algorithm string

const (
	// MergeSort will result in a more accurate result, but is slower
	MergeSort Algorithm = "mergesort"
	// Quicksort will result in a less accurate result, but is faster. Though the
	// results are mostly good enough
	QuickSelect Algorithm = "quickselect"
)

const (
	red = iota
	green
	blue
)

type point struct {
	X     int
	Y     int
	Red   uint32
	Green uint32
	Blue  uint32
}

func Quantize(src image.Image, palette color.Palette, algo Algorithm) (image.Image, error) {
	pixels := imageToBucket(src)
	buckets := make([][]point, len(palette))
	index := 0

	medianCut(pixels, buckets, &index, 0, len(palette)-1, algo)

	//colorPalette := make(color.Palette, len(pa))
	paletted := image.NewPaletted(src.Bounds(), palette)
	// calculate the average r,g,b for each bucket and use that average
	// at the color. Let's also fill in the paletted image
	for i, bucket := range buckets {
		var sumRed, sumGreen, sumBlue int64
		for _, p := range bucket {
			sumRed += int64(p.Red)
			sumGreen += int64(p.Green)
			sumBlue += int64(p.Blue)
		}

		l := int64(len(bucket))
		palette[i] = color.NRGBA64{
			R: uint16(sumRed / l),
			G: uint16(sumGreen / l),
			B: uint16(sumBlue / l),
			A: 0xffff,
		}

		for _, p := range bucket {
			paletted.SetColorIndex(p.X, p.Y, uint8(i))
		}
	}

	return paletted.SubImage(
		image.Rectangle{
			Min: image.Point{src.Bounds().Min.X, src.Bounds().Min.Y},
			Max: image.Point{src.Bounds().Max.X, src.Bounds().Max.Y},
		},
	), nil
}

func imageToBucket(img image.Image) []point {
	bounds := img.Bounds()
	// todo. I probably know the size here
	out := make([]point, 0)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			out = append(out, point{x, y, r, g, b})
		}
	}

	return out
}

func medianCut(pixels []point, buckets [][]point, index *int, depth, ncolors int, algo Algorithm) {
	// base case
	if float64(depth) >= math.Log2(float64(ncolors)) {
		buckets[*index] = make([]point, len(pixels))
		copy(buckets[*index], pixels)
		*index++
		return
	}

	// find the color (r,g,b) with the greatest range
	channel := greatestRange(pixels)

	if algo == MergeSort {
		// sort the pixels by that color
		pixels = sortByColorChannel(channel, pixels)
		// split the pixels and do it again
		i := len(pixels) / 2
		// left sublist
		medianCut(pixels[:i], buckets, index, depth+1, ncolors, algo)
		// right sublist
		medianCut(pixels[i:], buckets, index, depth+1, ncolors, algo)

	} else {
		medianIndex := len(pixels) / 2
		quickselect(pixels, 0, len(pixels)-1, medianIndex, channel)
		// left sublist
		medianCut(pixels[:medianIndex], buckets, index, depth+1, ncolors, algo)
		// right sublist
		medianCut(pixels[medianIndex:], buckets, index, depth+1, ncolors, algo)
	}

}

// greatestRange finds the channel (red, green or blue)
// with the greates range
func greatestRange(pts []point) int {
	var maxRed uint32
	var maxGreen uint32
	var maxBlue uint32

	minRed := uint32(math.MaxUint32)
	minGreen := uint32(math.MaxUint32)
	minBlue := uint32(math.MaxUint32)

	for i := 0; i < len(pts); i++ {
		pt := pts[i]

		if pt.Red < minRed {
			minRed = pt.Red
		}

		if pt.Red > maxRed {
			maxRed = pt.Red
		}

		if pt.Green < minGreen {
			minGreen = pt.Green
		}

		if pt.Green > maxGreen {
			maxGreen = pt.Green
		}

		if pt.Blue < minBlue {
			minBlue = pt.Blue
		}

		if pt.Blue > maxBlue {
			maxBlue = pt.Blue
		}
	}

	maxRange := red
	rg := maxRed - minRed
	if maxGreen-minGreen > rg {
		maxRange = green
		rg = maxGreen - minGreen
	}

	if maxBlue-minBlue > rg {
		maxRange = blue
	}

	return maxRange

}
