package container_with_most_water

import "testing"

func TestMaxArea(t *testing.T) {
	fixtures := []struct {
		name     string
		height   []int
		expected int
	}{
		{name: "test #1", height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expected: 49},
		{name: "test #2", height: []int{1, 1}, expected: 1},
	}

	for _, tt := range fixtures {
		got := MaxArea(tt.height)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
