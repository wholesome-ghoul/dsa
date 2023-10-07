package minimum_size_subarray_sum

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	t.Run("one element", func(t *testing.T) {
		got := MinSubArrayLen(7, []int{6})
		expected := 0

		assert(t, got, expected)

		got = MinSubArrayLen(7, []int{7})
		expected = 1

		assert(t, got, expected)
	})

	t.Run("repeated elements", func(t *testing.T) {
		got := MinSubArrayLen(4, []int{1, 4, 4})
		expected := 1

		assert(t, got, expected)
	})

	t.Run("duplicated values", func(t *testing.T) {
		got := MinSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1})
		expected := 0

		assert(t, got, expected)

		got = MinSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
		expected = 11

		assert(t, got, expected)
	})

	t.Run("no match returns 0", func(t *testing.T) {
		got := MinSubArrayLen(7, []int{6})
		expected := 0

		assert(t, got, expected)
	})

	t.Run("returns minimum array len", func(t *testing.T) {
		got := MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
		expected := 2

		assert(t, got, expected)
	})

	t.Run("success 1", func(t *testing.T) {
		got := MinSubArrayLen(11, []int{1, 2, 3, 4, 5})
		expected := 3

		assert(t, got, expected)
	})
}

func assert(t testing.TB, got, expected int) {
	t.Helper()

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
