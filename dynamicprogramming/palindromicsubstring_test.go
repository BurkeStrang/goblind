package dynamicprogramming

import "testing"

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abc", 3},       // "a", "b", "c"
		{"aaa", 6},       // "a", "a", "a", "aa", "aa", "aaa"
		{"racecar", 10},  // "r", "a", "c", "e", "c", "a", "r", "cec", "aceca", "racecar"
		{"", 0},          // No substrings
		{"a", 1},         // "a"
		{"ab", 2},        // "a", "b"
	}

	for _, test := range tests {
		result := CountSubstrings(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
