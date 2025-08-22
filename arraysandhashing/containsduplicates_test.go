package arraysandhashing

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{2, 2, 2, 2}, true},
	}

	for _, tt := range tests {
		result := ContainsDuplicate(tt.nums)
		if result != tt.expected {
			t.Errorf("ContainsDuplicate(%v) = %v; want %v", tt.nums, result, tt.expected)
		}
	}
}
