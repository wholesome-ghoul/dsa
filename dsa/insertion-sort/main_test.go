package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted array",
			input:    []int{4, 2, 1, 5, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element array",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := InsertionSort(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("%v: Expected %v, but got %v", test.name, test.expected, result)
			}
		})
	}
}
