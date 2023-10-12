package sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "array with length 1 returns self", input: []int{1}, expected: []int{1}},
		{name: "returns sorted array 1", input: []int{2, 0, 2, 1, 1, 0}, expected: []int{0, 0, 1, 1, 2, 2}},
		{name: "returns sorted array 2", input: []int{2, 0, 1}, expected: []int{0, 1, 2}},
	}

	for _, tt := range table {
		SortColors(tt.input)
		assert(t, tt.name, tt.expected, tt.input)
	}
}

func assert(t testing.TB, name string, expected, got []int) {
	t.Helper()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %#v, got %#v", expected, got)
	}
}
