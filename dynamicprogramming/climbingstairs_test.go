package dynamicprogramming

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{1, 1},
	}

	for _, tt := range tests {
		result := ClimbStairs(tt.n)
		if result != tt.expected {
			t.Errorf("ClimbStairs(%d) = %d; want %d", tt.n, result, tt.expected)
		}
	}
}
