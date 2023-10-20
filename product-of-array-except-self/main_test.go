package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	fixtures := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{name: "test #1", nums: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		// {name: "test #2", nums: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
		// {name: "test #3", nums: []int{0, 0}, expected: []int{0, 0}},
		// {name: "test #4", nums: []int{1, 0}, expected: []int{0, 1}},
	}

	for _, tt := range fixtures {
		got := ProductExceptSelf(tt.nums)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got []int) {
	t.Helper()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
