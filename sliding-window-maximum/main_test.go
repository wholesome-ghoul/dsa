package sliding_window_maximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	fixtures := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		// {name: "test #1", nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, expected: []int{3, 3, 5, 5, 6, 7}},
		// {name: "test #2", nums: []int{7, 2, 4}, k: 2, expected: []int{7, 4}},
		{name: "test #3", nums: []int{1, 3, 1, 2, 0, 5}, k: 3, expected: []int{3, 3, 2, 5}},
	}

	for _, tt := range fixtures {
		got := MaxSlidingWindow(tt.nums, tt.k)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got []int) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
