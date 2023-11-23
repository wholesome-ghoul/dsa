package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	fixtures := []struct {
		name     string
		s        string
		expected bool
	}{
		{name: "test #1", s: "race a car", expected: false},
		{name: "test #2", s: "A man, a plan, a canal: Panama", expected: true},
		{name: "test #3", s: " ", expected: true},
		{name: "test #4", s: "0P", expected: false},
		{name: "test #5", s: ".a", expected: true},
	}

	for _, tt := range fixtures {
		got := IsPalindrome(tt.s)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got bool) {
	t.Helper()
	if expected != got {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
