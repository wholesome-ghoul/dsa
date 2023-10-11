package minimum_window_substring

import "testing"

func TestMinWindow(t *testing.T) {
	table := []struct {
		name     string
		s        string
		t        string
		expected string
	}{
		{name: "s and t are the same #1", s: "a", t: "a", expected: "a"},
		{name: "s and t are the same #2", s: "abc", t: "abc", expected: "abc"},
		{name: "s length is less than t", s: "abc", t: "abcd", expected: ""},
		{name: "s does not have all characters from t #1", s: "abcefg", t: "abcd", expected: ""},
		{name: "s does not have all characters from t #1", s: "abcefg", t: "aa", expected: ""},
		{name: "s has all characters from t #1", s: "adobecodebanc", t: "abc", expected: "banc"},
		{name: "s has all characters from t #2", s: "bbaac", t: "aba", expected: "baa"},
	}

	for _, tt := range table {
		got := MinWindow(tt.s, tt.t)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name, expected, got string) {
	t.Helper()

	if got != expected {
		t.Errorf("%s: expected %s, got %s", name, expected, got)
	}
}
