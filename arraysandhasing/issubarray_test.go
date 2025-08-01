package arraysandhasing

import "testing"

func TestIsSubArray(t *testing.T) {
	tests := []struct {
		main    []int
		sub     []int
		want    bool
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3}, true},
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}, true},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, true},
		{[]int{1, 2, 3, 4, 5}, []int{5, 6}, false},
		{[]int{1, 2, 3, 4, 5}, []int{6}, false},
		{[]int{}, []int{1}, false},
	}

	for _, tt := range tests {
		got := IsSubArray(tt.main, tt.sub)
		if got != tt.want {
			t.Errorf("IsSubArray(%v, %v) = %v; want %v", tt.main, tt.sub, got, tt.want)
		}
	}
}
