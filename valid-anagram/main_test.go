package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	fixtures := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{name: "test #1", s: "anagram", t: "nagaram", expected: true},
		{name: "test #2", s: "car", t: "rat", expected: false},
	}

	for _, tt := range fixtures {
		got := IsAnagram(tt.s, tt.t)
		assert(t, tt.name, tt.expected, got)
	}
}

func assert(t testing.TB, name string, expected, got bool) {
	t.Helper()
	if expected != got {
		t.Errorf("%s: expected %#v, got %#v", name, expected, got)
	}
}
