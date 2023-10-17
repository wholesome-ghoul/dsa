package daily_temperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	fixtures := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "temperatures with one element returns [0]", input: []int{30}, expected: []int{0}},
		{name: "test #1", input: []int{30, 60, 90}, expected: []int{1, 1, 0}},
		{name: "test #2", input: []int{73, 74, 75, 71, 69, 72, 76, 73}, expected: []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{name: "test #3", input: []int{73, 72, 72, 76, 79}, expected: []int{3, 2, 1, 1, 0}},
	}

	for _, tt := range fixtures {
		got := DailyTemperatures(tt.input)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got []int) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
