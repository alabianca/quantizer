package quantizer

import (
	"image"
	"image/color"
	"math"
)

const (
	red = iota
	green
	blue
)

type Point struct {
	x     int
	y     int
	Red   uint32
	Green uint32
	Blue  uint32
}

func Quant(src image.Image, colors []Point) (image.Image, error) {
	pixels := imageToBucket(src)
	buckets := make([][]Point, len(colors))
	index := 0

	medianCut(pixels, buckets, &index, 0, len(colors)-1)

	colorPalette := make(color.Palette, len(colors))
	paletted := image.NewPaletted(src.Bounds(), colorPalette)
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
		colors[i] = Point{
			Red:   uint32((sumRed / l) >> 8),
			Green: uint32((sumGreen / l) >> 8),
			Blue:  uint32((sumBlue / l) >> 8),
		}

		colorPalette[i] = color.NRGBA64{
			R: uint16(sumRed / l),
			G: uint16(sumGreen / l),
			B: uint16(sumBlue / l),
			A: 0xffff,
		}

		for _, p := range bucket {
			paletted.SetColorIndex(p.x, p.y, uint8(i))
		}
	}

	return paletted.SubImage(
		image.Rectangle{
			Min: image.Point{src.Bounds().Min.X, src.Bounds().Min.Y},
			Max: image.Point{src.Bounds().Max.X, src.Bounds().Max.Y},
		},
	), nil
}

func imageToBucket(img image.Image) []Point {
	bounds := img.Bounds()
	// todo. I probably know the size here
	out := make([]Point, 0)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			out = append(out, Point{x, y, r, g, b})
		}
	}

	return out
}

func medianCut(pixels []Point, buckets [][]Point, index *int, depth, ncolors int) {
	// base case
	if float64(depth) >= math.Log2(float64(ncolors)) {
		buckets[*index] = make([]Point, len(pixels))
		copy(buckets[*index], pixels)
		*index++
		return
	}

	// find the color (r,g,b) with the greatest range
	rn := greatestRange(pixels)
	// sort the pixels by that color
	pixels = sortByColorChannel(rn, pixels)
	// split the pixels and do it again
	i := len(pixels) / 2
	// left sublist
	medianCut(pixels[:i], buckets, index, depth+1, ncolors)
	// right sublist
	medianCut(pixels[i:], buckets, index, depth+1, ncolors)

}

func greatestRange(pts []Point) int {
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

func sortByColorChannel(channel int, pts []Point) []Point {
	return mergeSort(pts, channel)
}

func mergeSort(items []Point, channel int) []Point {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]Point, middle)
		right = make([]Point, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left, channel), mergeSort(right, channel), channel)
}

func merge(left, right []Point, channel int) (result []Point) {
	result = make([]Point, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if channelFromPoint(channel, left[0]) < channelFromPoint(channel, right[0]) {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func channelFromPoint(channel int, p Point) uint32 {
	c := p.Red
	switch channel {
	case green:
		c = p.Green
	case blue:
		c = p.Blue
	}

	return c
}
