package first_missing_positive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	fixtures := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "test #1", nums: []int{1, 2, 0}, expected: 3},
		{name: "test #2", nums: []int{3, 4, -1, 1}, expected: 2},
		{name: "test #3", nums: []int{7, 8, 9, 11, 12}, expected: 1},
		{name: "test #4", nums: []int{1}, expected: 2},
		{name: "test #5", nums: []int{-5}, expected: 1},
		{name: "test #6", nums: []int{-1, -2}, expected: 1},
		{name: "test #7", nums: []int{1, 1}, expected: 2},
		{name: "test #8", nums: []int{-1, 4, 2, 1, 9, 10}, expected: 3},
	}

	for _, tt := range fixtures {
		got := FirstMissingPositive(tt.nums)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got int) {
	t.Helper()

	if expected != got {
		t.Errorf("%s: expected %d, got %d", name, expected, got)
	}
}
