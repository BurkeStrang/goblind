package arraysandhashing

import "testing"

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{}, 0},
		{[]int{1, 2, 0, 1}, 3},
		{[]int{10}, 1},
		{[]int{1, 9, 3, 10, 4, 20, 2}, 4},
	}

	for _, tt := range tests {
		result := LongestConsecutiveSequence(tt.nums)
		if result != tt.expected {
			t.Errorf("LongestConsecutiveSequence(%v) = %d; want %d", tt.nums, result, tt.expected)
		}
	}
}
