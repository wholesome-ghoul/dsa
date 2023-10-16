package merge_sorted_array

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	fixture := []struct {
		name     string
		expected []int
		nums1    []int
		m        int
		nums2    []int
		n        int
	}{
		{name: "nums2 is zero len array", nums1: []int{1}, nums2: []int{}, m: 1, n: 0, expected: []int{1}},
		{name: "nums1 is zero len array", nums1: []int{0}, nums2: []int{1}, m: 0, n: 1, expected: []int{1}},
		{name: "test #1", nums1: []int{1, 2, 3, 0, 0, 0}, nums2: []int{2, 5, 6}, m: 3, n: 3, expected: []int{1, 2, 2, 3, 5, 6}},
	}

	for _, tt := range fixture {
		Merge(tt.nums1, tt.m, tt.nums2, tt.n)
		assert(t, tt.name, tt.expected, tt.nums1, tt.nums2, tt.m, tt.n)
	}
}

func assert(t testing.TB, name string, expected, nums1, nums2 []int, m, n int) {
	t.Helper()
	if !reflect.DeepEqual(expected, nums1) {
		t.Errorf("%s: expected %#v, got %#v", name, expected, nums1)
	}
}
