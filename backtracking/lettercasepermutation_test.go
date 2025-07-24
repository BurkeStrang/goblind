package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"Ad2", []string{"Ad2", "aD2", "ad2", "AD2"}},
		{"a1b2", []string{"a1b2", "A1b2", "a1B2", "A1B2"}},
		{"123", []string{"123"}},
		{"", []string{""}},
	}

	for _, tt := range tests {
		got := LetterCasePermutation(tt.input)
		sort.Strings(got)
		sort.Strings(tt.expected)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("LetterCasePermutation(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}
