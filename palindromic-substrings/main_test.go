package palindromic_substrings

import "testing"

func TestCountSubstrings(t *testing.T) {
	table := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "single character string", input: "a", expected: 1},
		{name: "all letters are same", input: "aaa", expected: 6},
		{name: "test #1", input: "abc", expected: 3},
	}

	for _, tt := range table {
		got := CountSubstrings(tt.input)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got int) {
	t.Helper()

	if got != expected {
		t.Errorf("%s: expected %d, got %d", name, expected, got)
	}
}
