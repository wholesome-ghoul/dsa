package longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	table := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "empty string returns 0", input: "", expected: 0},
		{name: "string with length 1 returns 1", input: "a", expected: 1},
		{name: "repeated chars in string returns 1", input: "aaaaa", expected: 1},
		{name: "string with duplicate chars #1", input: "abcabcbb", expected: 3},
		{name: "string with duplicate chars #2", input: "pwwkew", expected: 3},
		{name: "string with duplicate chars #3", input: "xxzqi", expected: 4},
		{name: "string with duplicate chars #4", input: "ckilbkd", expected: 5},
		{name: "string with duplicate chars #5", input: "tmmzuxt", expected: 5},
		{name: "string with duplicate chars #6", input: "ggububgvfk", expected: 6},
	}

	for _, tt := range table {
		got := lengthOfLongestSubstring(tt.input)
		assert(t, tt.name, got, tt.expected)
	}
}

func assert(t testing.TB, name string, got, expected int) {
	t.Helper()

	if got != expected {
		t.Errorf("%s: got %d, expected %d", name, got, expected)
	}
}
