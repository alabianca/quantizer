package mediancut

import "testing"

func TestQuickSelect(t *testing.T) {
	var image = [][]Point{
		{{0, 0, 120, 200, 150}, {0, 0, 10, 50, 255}, {0, 0, 200, 0, 100}},
		{{0, 0, 15, 25, 35}, {0, 0, 75, 85, 95}, {0, 0, 135, 145, 155}},
		{{0, 0, 5, 10, 15}, {0, 0, 25, 35, 45}, {0, 0, 55, 65, 75}},
	}

	points := flatten(image)
	if len(points) != 9 {
		t.Fatalf("expected length %d but got %d", 9, len(points))
	}

	//Red channel sorted 5, 10, 15, 25, 55, 75, 120, 135, 200
	//Red channel unsorted 120, 10, 200, 15, 75, 135, 5, 25, 55
	k := len(points) / 2 // 75 (index 4)
	quickselect(points, 0, len(points)-1, k, red)
	if points[4].Red != 55 {
		t.Fatalf("expected pivot %d to be at index %d", 55, 4)
	}
}
