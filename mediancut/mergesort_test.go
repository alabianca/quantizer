package mediancut

import "testing"

func TestMergeSort(t *testing.T) {
	// Point{X,Y,R,G,B}
	var image = [][]Point{
		{{0, 0, 120, 200, 150}, {0, 0, 10, 50, 255}, {0, 0, 200, 0, 100}},
		{{0, 0, 15, 25, 35}, {0, 0, 75, 85, 95}, {0, 0, 135, 145, 155}},
		{{0, 0, 5, 10, 15}, {0, 0, 25, 35, 45}, {0, 0, 55, 65, 75}},
	}

	points := flatten(image)
	if len(points) != 9 {
		t.Fatalf("expected length %d but got %d", 9, len(points))
	}
	sorted := sortByColorChannel(red, points)
	expected := []Point{
		{0, 0, 5, 10, 15},
		{0, 0, 10, 50, 255},
		{0, 0, 15, 25, 35},
		{0, 0, 25, 35, 45},
		{0, 0, 55, 65, 75},
		{0, 0, 75, 85, 95},
		{0, 0, 120, 200, 150},
		{0, 0, 135, 145, 155},
		{0, 0, 200, 0, 100},
	}

	for i, p := range sorted {
		if p.Red != expected[i].Red {
			t.Fatalf("For %dth element expected red to be %d, but got %d\n", i+1, expected[i].Red, p.Red)
		}
	}
}

func flatten(img [][]Point) []Point {
	var flattened []Point
	for _, row := range img {
		flattened = append(flattened, row...)
	}
	return flattened
}
