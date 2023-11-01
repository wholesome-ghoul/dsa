package search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	fixtures := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{name: "test #1", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, expected: 4},
		{name: "test #2", nums: []int{5, 1, 3}, target: 5, expected: 0},
	}

	for _, tt := range fixtures {
		got := Search(tt.nums, tt.target)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
