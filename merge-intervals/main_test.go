package merge_intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	fixtures := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{name: "intervals with one element returns self", intervals: [][]int{{1, 2}}, expected: [][]int{{1, 2}}},
		{name: "test #1", intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, expected: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{name: "test #2", intervals: [][]int{{1, 4}, {0, 4}}, expected: [][]int{{0, 4}}},
	}

	for _, tt := range fixtures {
		got := Merge(tt.intervals)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got [][]int) {
	t.Helper()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
