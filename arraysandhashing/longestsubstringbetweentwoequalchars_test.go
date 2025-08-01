package arraysandhashing

import "testing"

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	tests := []struct {
		s      string
		expect int
	}{
		{"aa", 0},
		{"abca", 2},
		{"cbzxy", -1},
		{"aabb", 0},
		{"abcabc", 2},
		{"a", -1},
		{"abccba", 4},
	}

	for _, tt := range tests {
		got := MaxLengthBetweenEqualCharacters(tt.s)
		if got != tt.expect {
			t.Errorf("MaxLengthBetweenEqualCharacters(%q) = %d; want %d", tt.s, got, tt.expect)
		}
	}
}
