package non_overlapping_intervals

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	fixtures := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{name: "test #1", intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, expected: 1},
		{name: "test #2", intervals: [][]int{{1, 2}, {1, 2}, {1, 2}}, expected: 2},
		{name: "test #3", intervals: [][]int{{1, 2}, {2, 3}}, expected: 0},
		{name: "test #4", intervals: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}, expected: 2},
	}

	for _, tt := range fixtures {
		got := EraseOverlapIntervals(tt.intervals)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got int) {
	t.Helper()

	if expected != got {
		t.Errorf("%s: expected %d, got %d", name, expected, got)
	}
}
