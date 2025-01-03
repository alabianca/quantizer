package mediancut

func sortByColorChannel(channel int, pts []Point) []Point {
	return mergeSort(pts, channel)
}

func mergeSort(items []Point, channel int) []Point {
	var ln = len(items)

	if ln == 1 {
		return items
	}

	middle := int(ln / 2)
	left := items[:middle]
	right := items[middle:]

	return merge(mergeSort(left, channel), mergeSort(right, channel), channel)
}

func merge(left, right []Point, channel int) (result []Point) {
	result = make([]Point, len(left)+len(right))

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if getChannelFromPoint(channel, left[i]) < getChannelFromPoint(channel, right[j]) {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// Copy remaining elements from `left`
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	// Copy remaining elements from `right`
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return
}

// getChannelFromPoint returns the red,green or blue channel from Point
func getChannelFromPoint(channel int, p Point) uint32 {
	c := p.Red
	switch channel {
	case green:
		c = p.Green
	case blue:
		c = p.Blue
	}

	return c
}
