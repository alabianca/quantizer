package mediancut

func sortByColorChannel(channel int, pts []point) []point {
	return mergeSort(pts, channel)
}

func mergeSort(items []point, channel int) []point {
	var ln = len(items)

	if ln == 1 {
		return items
	}

	middle := int(ln / 2)
	left := items[:middle]
	right := items[middle:]

	return merge(mergeSort(left, channel), mergeSort(right, channel), channel)
}

func merge(left, right []point, channel int) (result []point) {
	result = make([]point, len(left)+len(right))

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
func getChannelFromPoint(channel int, p point) uint32 {
	c := p.Red
	switch channel {
	case green:
		c = p.Green
	case blue:
		c = p.Blue
	}

	return c
}
