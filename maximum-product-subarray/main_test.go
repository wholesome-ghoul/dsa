package maximum_product_subarray

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	fixtures := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "test #1", nums: []int{2, 3, -2, 4}, expected: 6},
		{name: "test #2", nums: []int{-2, 0, -1}, expected: 0},
		{name: "test #3", nums: []int{-1, 3, 3, -1}, expected: 9},
		{name: "test #4", nums: []int{-3, 0, 1, -2}, expected: 1},
		{name: "test #5", nums: []int{-1, 4, -4, 5, -2, -1, -1, -2, -3}, expected: 960},
		{name: "test #6", nums: []int{-2}, expected: -2},
		{name: "test #7", nums: []int{-1, -2, -3, 0}, expected: 6},
	}

	for _, tt := range fixtures {
		got := MaxProduct(tt.nums)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
