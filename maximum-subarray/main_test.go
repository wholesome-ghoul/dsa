package maximum_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	fixtures := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "test #1", nums: []int{1}, expected: 1},
		{name: "test #2", nums: []int{5, 4, -1, 7, 8}, expected: 23},
		{name: "test #3", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6},
	}

	for _, tt := range fixtures {
		got := MaxSubArray(tt.nums)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
