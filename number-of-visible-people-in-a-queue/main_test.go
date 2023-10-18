package number_of_visible_people_in_a_queue

import (
	"reflect"
	"testing"
)

func TestCanSeePersonsCount(t *testing.T) {
	fixtures := []struct {
		name     string
		heights  []int
		expected []int
	}{
		{name: "heights with length one returns [0]", heights: []int{4}, expected: []int{0}},
		{name: "test #1", heights: []int{10, 6, 8, 5, 11, 9}, expected: []int{3, 1, 2, 1, 1, 0}},
		{name: "test #2", heights: []int{5, 1, 2, 3, 10}, expected: []int{4, 1, 1, 1, 0}},
	}

	for _, tt := range fixtures {
		got := CanSeePersonsCount(tt.heights)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got []int) {
	t.Helper()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
