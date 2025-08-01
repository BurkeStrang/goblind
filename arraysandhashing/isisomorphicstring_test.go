package arraysandhashing

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s, t     string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"ab", "aa", false},
		{"", "", true},
		{"a", "a", true},
		{"abc", "def", true},
		{"abc", "dee", false},
		{"abab", "baba", true},
		{"abab", "bbaa", false},
	}

	for _, tt := range tests {
		result := IsIsomorphic(tt.s, tt.t)
		if result != tt.expected {
			t.Errorf("IsIsomorphic(%q, %q) = %v; want %v", tt.s, tt.t, result, tt.expected)
		}
	}
}
