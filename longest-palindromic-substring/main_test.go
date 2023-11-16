package longest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	fixtures := []struct {
		name     string
		s        string
		expected string
	}{
		{name: "test #1", s: "babad", expected: "bab"},
		{name: "test #2", s: "cbbd", expected: "bb"},
		{name: "test #3", s: "abcdcbaqr", expected: "abcdcba"},
		{name: "test #4", s: "aaaa", expected: "aaaa"},
	}

	for _, tt := range fixtures {
		got := LongestPalindrome(tt.s)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got string) {
	t.Helper()
	if expected != got {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
