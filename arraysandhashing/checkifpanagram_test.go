package arraysandhashing

import "testing"

func TestCheckIfPanagram(t *testing.T) {
	tests := []struct {
		sentence string
		want     bool
	}{
		{"The quick brown fox jumps over the lazy dog", true},
		{"abcdefghijklmnopqrstuvwxyz", true},
		{"Hello World", false},
		{"", false},
		{"Pack my box with five dozen liquor jugs", true},
		{"Sphinx of black quartz, judge my vow", true},
		{"abcde", false},
	}

	for _, tt := range tests {
		got := CheckIfPanagram(tt.sentence)
		if got != tt.want {
			t.Errorf("checkIfPanagram(%q) = %v, want %v", tt.sentence, got, tt.want)
		}
	}
}
