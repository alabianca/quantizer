package mediancut

func partition(points []Point, channel int, left, right int) int {
	pivot := getChannelFromPoint(channel, points[right]) // Use the rightmost element as the pivot
	i := left - 1

	for j := left; j < right; j++ {
		if getChannelFromPoint(channel, points[j]) < pivot {
			i++
			points[i], points[j] = points[j], points[i] // Swap smaller elements to the left
		}
	}
	points[i+1], points[right] = points[right], points[i+1] // Place pivot in correct position
	return i + 1
}

func quickselect(points []Point, left, right, k int, channel int) {
	if left >= right {
		return
	}

	pivotIndex := partition(points, channel, left, right)
	if k == pivotIndex {
		return
	} else if k < pivotIndex {
		quickselect(points, left, pivotIndex-1, k, channel)
	} else {
		quickselect(points, pivotIndex+1, right, k, channel)
	}
}
