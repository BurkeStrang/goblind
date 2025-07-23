package backtracking

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, tt := range tests {
		result := CombinationSum(tt.candidates, tt.target)
		if !equalUnordered(result, tt.expected) {
			t.Errorf("CombinationSum(%v, %d) = %v; want %v", tt.candidates, tt.target, result, tt.expected)
		}
	}
}

// Helper to compare two slices of slices regardless of order
func equalUnordered(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, x := range a {
		found := false
		for j, y := range b {
			if !used[j] && reflect.DeepEqual(x, y) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
