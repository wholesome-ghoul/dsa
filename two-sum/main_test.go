package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	fixtures := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{name: "test #1", nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{name: "test #2", nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{name: "test #3", nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}

	for _, tt := range fixtures {
		got := TwoSum(tt.nums, tt.target)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got []int) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
